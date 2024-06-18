package usecase

import (
	"errors"
	_ "foodOrder/domain/entity"
	"foodOrder/domain/model"
	"foodOrder/internal/api/payment/repository"

	_ "github.com/oklog/ulid/v2"
)

type PaymentUsecase struct {
	paymentRepo *repository.PaymentRepo
}

func NewPaymentUsecase(repo *repository.PaymentRepo) *PaymentUsecase {
	return &PaymentUsecase{paymentRepo: repo}
}

func (u *PaymentUsecase) CreatePayment(payment *model.CreatePayment) (model.Bill, error) {
	order, err := u.paymentRepo.GetOrderByTableNo(payment.TableNo)
	if err != nil {
		return model.Bill{}, err
	}

	if len(order) == 0 {
		return model.Bill{}, errors.New("no order found")
	}

	var total float64
	var bill model.Bill
	for _, v := range order {
		if v.Status == "done" || v.Status == "end" {
			food := u.paymentRepo.GetFoodByID(v.FoodId)
			bill.Detail = append(bill.Detail, model.BillDetail{
				FoodName: food.Name,
				Price:    float64(food.Price),
				Quantity: v.Quantity,
			})
			total += float64(v.Quantity) * float64(food.Price)
		}
	}

	bill.Total = total
	bill.TableNo = payment.TableNo
	bill.PreferenceID, err = u.paymentRepo.GetPreferenceID(payment.TableNo)
	if err != nil {
		return model.Bill{}, err
	}
	
	return bill, nil
}