package usecase

import (
	"foodOrder/domain/entity"
	"foodOrder/domain/model"
	"foodOrder/internal/api/restaurant/repository"
)

type RestaurantUsecase struct {
	restaurantRepo *repository.RestRepo
}

func NewRestaurantUsecase(repo *repository.RestRepo) *RestaurantUsecase {
	return &RestaurantUsecase{restaurantRepo: repo}
}

func (r *RestaurantUsecase) CreateRestaurant(restaurant *model.CreateRestaurant) error {
	dbRest := &entity.Restaurant{
		Name:       restaurant.Name,
		TotalTable: restaurant.TableNo,
	}

	err := r.restaurantRepo.CreateRestaurant(dbRest)
	if err != nil {
		return err
	}

	return nil
}

func (r *RestaurantUsecase) AdjustTable(restaurant *model.AdjustTable) error {
	dbRest := &entity.Restaurant{
		Name: 	 restaurant.Name,
		TotalTable: restaurant.TableNo,
	}

	err := r.restaurantRepo.AdjustTable(dbRest)
	if err != nil {
		return err
	}

	return nil
}