package usecase

import (
	"errors"
	"fmt"
	"foodOrder/domain/entity"
	"foodOrder/domain/model"
	"foodOrder/internal/api/restaurant/repository"

	"github.com/oklog/ulid/v2"
)

type RestaurantUsecase struct {
	restaurantRepo *repository.RestRepo
}

func NewRestaurantUsecase(repo *repository.RestRepo) *RestaurantUsecase {
	return &RestaurantUsecase{restaurantRepo: repo}
}

func (u *RestaurantUsecase) InitialTable(table *model.InitialTable) error {
	amount := table.TableNo
	fmt.Println(amount)
	for i := 1; i <= int(amount); i++ {
		dbTable := &entity.Table{
			TableNo: uint8(i),
		}

		fmt.Println(dbTable)

		err := u.restaurantRepo.InitialTable(dbTable)
		if err != nil {
			return errors.New("failed to initial table")
		}
	}

	return nil
}

func (u *RestaurantUsecase) GetAllTable() ([]model.TableDetail, error) {
	tables, err := u.restaurantRepo.GetALLTable()
	if err != nil {
		return nil, err
	}

	var tableDetail []model.TableDetail
	for _, v := range tables {
		tableDetail = append(tableDetail, model.TableDetail{
			TableNo: v.TableNo,
			Status: v.Status,
		})
	}

	return tableDetail, nil
}

func (u *RestaurantUsecase) GiveCustomerTable(table *model.GiveTable) (ulid.ULID, error) {
	tableNo := table.TableNo

	status, err := u.restaurantRepo.GetTable(tableNo)
	if err != nil {
		return ulid.ULID{}, errors.New("table not found")
	}

	if status.Status == "occupied" {
		return ulid.ULID{}, errors.New("table is occupied")
	}

	preferenceID, err := u.restaurantRepo.GiveCustomerTable(tableNo)
	if err != nil {
		return ulid.ULID{}, errors.New("failed to give table")
	}

	return preferenceID, nil

}