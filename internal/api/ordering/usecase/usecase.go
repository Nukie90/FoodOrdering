package usecase

import (
	"errors"
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
		TableNo: cart.TableNo,
		UserOrder: GuestID,
		FoodId: foodId,
		Quantity: cart.Quantity,
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
			TableNo: v.TableNo,
			FoodName: foodName,
			Quantity: v.Quantity,
		})
	}

	return cartDetail, nil
}

func (u *OrderingUsecase) SubmitCart(tableNo uint8)  error {
	cart, err := u.orderingRepo.CartDetail(tableNo)
	if err != nil {
		return err
	}

	var orderDetail []entity.Order
	for _, v := range cart {


		orderDetail = append(orderDetail, entity.Order{
			TableNo: v.TableNo,
			FoodId: v.FoodId,
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