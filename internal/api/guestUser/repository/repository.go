package repository

import (
	"foodOrder/domain/entity"

	"gorm.io/gorm"
)

type GuestRepo struct {
	db *gorm.DB
}

func NewGuestRepo(db *gorm.DB) *GuestRepo {
	return &GuestRepo{db: db}
}

func (r *GuestRepo) TableAmount() (uint8, error) {
	dbTx := r.db.Begin()
	defer dbTx.Rollback()

	var tableAmount uint8
	if err := dbTx.Model(&entity.Restaurant{}).Select("total_table").First(&tableAmount).Error; err != nil {
		return 0, err
	}

	return tableAmount, nil
}