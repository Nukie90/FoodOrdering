package repository

import (
	"errors"
	"foodOrder/domain/entity"
	"foodOrder/domain/model"

	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) CreateOrder(order *entity.Order) error {
	err := r.db.Create(order).Error
	if err != nil {
		return errors.New("failed to create order")
	}

	return nil
}