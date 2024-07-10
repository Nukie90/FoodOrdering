package usecase

import (
	"errors"
	"foodOrder/domain/entity"
	"foodOrder/domain/model"
	"foodOrder/internal/api/payment/repository"

	"github.com/oklog/ulid/v2"
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

func (u *PaymentUsecase) PayBill(preferenceID string, payment *model.PayBill) (interface{}, model.Bill, error) {
	tableNo := u.paymentRepo.GetTableNo(preferenceID)

	var receiptID ulid.ULID
	ulidPreferenceID, err := ulid.Parse(preferenceID)
	if err != nil {
		return nil, model.Bill{}, err
	}

	bill, err := u.CreatePayment(&model.CreatePayment{
		TableNo: uint(tableNo),
	})
	if err != nil {
		return 	nil, model.Bill{}, err
	}

	switch payment.PaymentMethod {
	case "cash":
		if payment.Amount < bill.Total {
			return nil, model.Bill{}, errors.New("insufficient amount")
		} else if payment.Amount > bill.Total {
			change := payment.Amount - bill.Total

			receiptID, err = u.paymentRepo.CreateBill(&entity.Payment{
				PreferenceID: ulidPreferenceID,
				Total:        bill.Total,
			})
			if err != nil {
				return nil, model.Bill{}, err
			}
			bill.ReceiptID = receiptID

			return change, bill, nil
		}
	case "credit":
		receiptID, err = u.paymentRepo.CreateBill(&entity.Payment{
			PreferenceID: ulidPreferenceID,
			Total:        bill.Total,
		})
		bill.ReceiptID = receiptID
		if err != nil {
			return nil, bill, err
		}
	default:
		return nil, model.Bill{}, errors.New("invalid payment method")
	}

	return nil, bill, nil
}
