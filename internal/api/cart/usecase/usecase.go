package usecase

import (
	"errors"
	"foodOrder/domain/entity"
	"foodOrder/domain/model"
	"foodOrder/internal/api/cart/repository"

	"github.com/oklog/ulid/v2"
)

type CartUsecase struct {
	cartRepo *repository.CartRepo
}

func NewCartUsecase(repo *repository.CartRepo) *CartUsecase {
	return &CartUsecase{cartRepo: repo}
}

func (u *CartUsecase) AddToCart(cart *model.AddToCart) error {
	GuestID, err := ulid.Parse(cart.UserOrder)
	if err != nil {
		return errors.New("failed to parse guest id")
	}

	foodId, err := u.cartRepo.GetFoodIdByName(cart.FoodName)
	if err != nil {
		return errors.New("failed to get food id")
	}

	dbCart := &entity.Cart{
		TableNo: cart.TableNo,
		UserOrder: GuestID,
		FoodId: foodId,
		Quantity: cart.Quantity,
	}

	err = u.cartRepo.AddToCart(dbCart)
	if err != nil {
		return err
	}

	return nil
}