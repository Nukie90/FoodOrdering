package repository

import (
	"foodOrder/domain/entity"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type PaymentRepo struct {
	Paymentdb *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{Paymentdb: db}
}


func (r *PaymentRepo) GetOrderByTableNo(tableNo uint) ([]entity.Order, error) {
	var orders []entity.Order
	err := r.Paymentdb.Where("table_no = ?", tableNo).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *PaymentRepo) GetFoodByID(foodId uint) entity.Food {
	dbTx := r.Paymentdb.Begin()
	defer dbTx.Rollback()

	var food entity.Food
	if err := dbTx.Where("food_id = ?", foodId).First(&food).Error; err != nil {
		return entity.Food{}
	}

	return food
}

func (r *PaymentRepo) GetPreferenceID(tableNo uint) (ulid.ULID, error) {
	dbTx := r.Paymentdb.Begin()
	defer dbTx.Rollback()

	var preference entity.TablePreference
	if err := dbTx.Where("table_no = ?", tableNo).First(&preference).Error; err != nil {
		return ulid.ULID{}, err
	}

	return preference.PreferenceID, nil
}
