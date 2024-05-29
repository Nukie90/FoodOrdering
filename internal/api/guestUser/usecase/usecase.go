package usecase

import (
	"errors"
	"fmt"
	"foodOrder/domain/model"
	"foodOrder/internal/api/guestUser/repository"
)

type GuestUsecase struct {
	guestRepo *repository.GuestRepo
}

func NewGuestUsecase(repo *repository.GuestRepo) *GuestUsecase {
	return &GuestUsecase{guestRepo: repo}
}

func (u *GuestUsecase) EnterTable(table *model.EnterTable) error {
	tableAmount, err := u.guestRepo.TableAmount()
	if err != nil {
		return err
	}
	fmt.Println(table.TableNo)

	if table.TableNo > tableAmount {
		return errors.New("table number is invalid")
	}

	return nil
}

func (u *GuestUsecase) SameTableGuest(table *model.EnterTable) error {
	return nil
}