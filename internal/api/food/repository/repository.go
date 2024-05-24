package repository

import (
	"errors"
	"foodOrder/domain/entity"
	"foodOrder/domain/model"

	"gorm.io/gorm"
)

type FoodRepo struct {
	foodDb *gorm.DB
}

func NewFoodRepo(db *gorm.DB) *FoodRepo {
	return &FoodRepo{foodDb: db}
}

func (f *FoodRepo) CreateFood(food *entity.Food) error {
	dbTx := f.foodDb.Begin()
	defer dbTx.Rollback()

	if err := dbTx.Create(food).Error; err != nil {
		return errors.New("failed to create food")
	}

	return dbTx.Commit().Error
}

func (f *FoodRepo) GetAllFoods() ([]model.FoodDetail, error) {
	var foods []entity.Food
	if err := f.foodDb.Find(&foods).Error; err != nil {
		return nil, err
	}

	var foodDetail []model.FoodDetail
	for _, food := range foods {
		foodDetail = append(foodDetail, model.FoodDetail{
			ID:          food.FoodId,
			Name:        food.Name,
			Description: food.Description,
			Price:       food.Price,
		})
	}

	return foodDetail, nil
}