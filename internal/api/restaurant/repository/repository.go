package repository

import (
	"foodOrder/domain/entity"

	"gorm.io/gorm"
)

type RestRepo struct {
	restDb *gorm.DB
}

func NewRestRepo(db *gorm.DB) *RestRepo {
	return &RestRepo{restDb: db}
}

func (r *RestRepo) CreateRestaurant(restaurant *entity.Restaurant) error {
	dbTx := r.restDb.Begin()
	defer dbTx.Rollback()

	if err := dbTx.Create(restaurant).Error; err != nil {
		return err
	}

	return dbTx.Commit().Error
}

func (r *RestRepo) AdjustTable(restaurant *entity.Restaurant) error {
	dbTx := r.restDb.Begin()
	defer dbTx.Rollback()

	if err := dbTx.Save(restaurant).Error; err != nil {
		return err
	}

	return dbTx.Commit().Error
}