package usecase

import (
	"foodOrder/domain/entity"
	"foodOrder/domain/model"
	"foodOrder/internal/api/ordering/repository"
)

type OrderingUsecase struct {
	orderRepo *repository.OrderRepo
}

func NewOrderingUsecase(repo *repository.OrderRepo) *OrderingUsecase {
	return &OrderingUsecase{orderRepo: repo}
}

func (u *OrderingUsecase) CreateOrder(order *entity.Order) error {
	err := u.orderRepo.CreateOrder(order)
	if err != nil {
		return err
	}

	return nil
}