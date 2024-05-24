package usecase

import (
	"foodOrder/domain/entity"
	"foodOrder/domain/model"
	"foodOrder/internal/api/food/repository"
)

type FoodUsecase struct {
	foodRepo *repository.FoodRepo
}

func NewFoodUsecase(repo *repository.FoodRepo) *FoodUsecase {
	return &FoodUsecase{foodRepo: repo}
}


func (f *FoodUsecase)CreateFood(food *model.CreateFood) error {
	dbFood := &entity.Food{
		Name: food.Name,
		Description: food.Description,
		Price: food.Price,
	}

	err := f.foodRepo.CreateFood(dbFood)
	if err != nil {
		return err
	}

	return nil
}

func (f *FoodUsecase)GetAllFoods() ([]model.FoodDetail, error) {
	foods, err := f.foodRepo.GetAllFoods()
	if err != nil {
		return nil, err
	}

	return foods, nil
}
