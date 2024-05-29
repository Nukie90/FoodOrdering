package repository

import (
	"errors"
	"foodOrder/domain/entity"

	"gorm.io/gorm"
)

type CartRepo struct {
	db *gorm.DB
}

func NewCartRepo(db *gorm.DB) *CartRepo {
	return &CartRepo{db: db}
}

func (r *CartRepo) AddToCart(cart *entity.Cart) error {
	err := r.db.Create(cart).Error
	if err != nil {
		return errors.New("failed to add to cart")
	}

	return nil
}

func (r *CartRepo) GetFoodIdByName(name string) (uint, error) {
	dbTx := r.db.Begin()
	defer dbTx.Rollback()

	var food entity.Food
	if err := dbTx.Where("name = ?", name).First(&food).Error; err != nil {
		return 0, err
	}

	return food.FoodId, nil
}