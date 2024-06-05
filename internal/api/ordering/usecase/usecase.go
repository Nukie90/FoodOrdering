package usecase

import (
	"errors"
	"fmt"
	"foodOrder/domain/entity"
	"foodOrder/domain/model"
	"foodOrder/internal/api/ordering/repository"

	"github.com/oklog/ulid/v2"
)

type OrderingUsecase struct {
	orderingRepo *repository.OrderingRepo
}

func NewOrderingUsecase(repo *repository.OrderingRepo) *OrderingUsecase {
	return &OrderingUsecase{orderingRepo: repo}
}

func (u *OrderingUsecase) AddToCart(cart *model.AddToCart) error {
	GuestID, err := ulid.Parse(cart.UserOrder)
	if err != nil {
		return errors.New("failed to parse guest id")
	}

	foodId, err := u.orderingRepo.GetFoodIdByName(cart.FoodName)
	if err != nil {
		return errors.New("failed to get food id")
	}

	dbCart := &entity.Cart{
		TableNo:   cart.TableNo,
		UserOrder: GuestID,
		FoodId:    foodId,
		Quantity:  cart.Quantity,
	}

	err = u.orderingRepo.AddToCart(dbCart)
	if err != nil {
		return err
	}

	return nil
}

func (u *OrderingUsecase) CartDetail(tableNo uint8) ([]model.CartDetail, error) {
	cart, err := u.orderingRepo.CartDetail(tableNo)
	if err != nil {
		return nil, err
	}

	var cartDetail []model.CartDetail
	for _, v := range cart {
		foodName, err := u.orderingRepo.GetFoodNameById(v.FoodId)
		if err != nil {
			return nil, err
		}

		cartDetail = append(cartDetail, model.CartDetail{
			TableNo:  v.TableNo,
			FoodName: foodName,
			Quantity: v.Quantity,
		})
	}

	return cartDetail, nil
}

func (u *OrderingUsecase) SubmitCart(tableNo uint8) error {
	cart, err := u.orderingRepo.CartDetail(tableNo)
	if err != nil {
		return err
	}

	var orderDetail []entity.Order
	for _, v := range cart {
		orderDetail = append(orderDetail, entity.Order{
			TableNo:  v.TableNo,
			FoodId:   v.FoodId,
			Quantity: v.Quantity,
		})
	}

	err = u.orderingRepo.SubmitCart(orderDetail)
	if err != nil {
		return err
	}

	err = u.orderingRepo.DeleteCart(tableNo)
	if err != nil {
		return err
	}

	return nil
}

func (u *OrderingUsecase) ReceiveOrder() ([]model.TableOrder, error) {
	totalTable, err := u.orderingRepo.TableAmount()
	if err != nil {
		return nil, err
	}

	var tableOrder []model.TableOrder
	for i := 1; i <= int(totalTable); i++ {
		order, err := u.orderingRepo.GetOrder(uint8(i))
		if err != nil {
			return nil, err
		}

		var orderDetail []model.OrderDetail
		for _, v := range order {
			foodName, err := u.orderingRepo.GetFoodNameById(v.FoodId)
			if err != nil {
				return nil, err
			}

			orderDetail = append(orderDetail, model.OrderDetail{
				OderId:   v.OrderId,
				FoodName: foodName,
				Quantity: v.Quantity,
				Status:   v.Status,
			})
		}

		tableOrder = append(tableOrder, model.TableOrder{
			TableNo: i,
			Detail:  orderDetail,
		})
	}

	fmt.Println(tableOrder)

	return tableOrder, nil
}

func (u *OrderingUsecase) SendRobot(reqForm *model.SendRobotRequest) (uint8, error) {
	order, err := u.orderingRepo.GetOrderByID(reqForm.OrderID)
	if err != nil {
		return 0, err
	}

	for _, v := range order {
		v.Status = "done"
		err = u.orderingRepo.UpdateOrder(v)
		if err != nil {
			return 0, err
		}
	}

	return order[0].TableNo, nil
}

func (u *OrderingUsecase) ReceiveRobot(tableNo uint8) error {
	order, err := u.orderingRepo.GetOrder(tableNo)
	if err != nil {
		return err
	}

	for _, v := range order {
		if v.Status == "done" {
			v.Status = "end"
			err = u.orderingRepo.UpdateOrder(v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
