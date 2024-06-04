package repository

import (
	"errors"
	"foodOrder/domain/entity"

	"gorm.io/gorm"
)

type OrderingRepo struct {
	db *gorm.DB
}

func NewOrderingRepo(db *gorm.DB) *OrderingRepo {
	return &OrderingRepo{db: db}
}

func (r *OrderingRepo) AddToCart(cart *entity.Cart) error {
	err := r.db.Create(cart).Error
	if err != nil {
		return errors.New("failed to add to cart")
	}

	return nil
}

func (r *OrderingRepo) GetFoodIdByName(name string) (uint, error) {
	dbTx := r.db.Begin()
	defer dbTx.Rollback()

	var food entity.Food
	if err := dbTx.Where("name = ?", name).First(&food).Error; err != nil {
		return 0, err
	}

	return food.FoodId, nil
}

func (r *OrderingRepo) CartDetail(tableNo uint8) ([]entity.Cart, error) {
	var cart []entity.Cart
	if err := r.db.Where("table_no = ?", tableNo).Find(&cart).Error; err != nil {
		return nil, err
	}

	return cart, nil
}

func (r *OrderingRepo) GetFoodNameById(foodId uint) (string, error) {
	dbTx := r.db.Begin()
	defer dbTx.Rollback()

	var food entity.Food
	if err := dbTx.Where("food_id = ?", foodId).First(&food).Error; err != nil {
		return "", err
	}

	return food.Name, nil
}

func (r *OrderingRepo) DeleteCart(tableNo uint8) error {
	dbTx := r.db.Begin()
	defer dbTx.Rollback()

	if err := dbTx.Where("table_no = ?", tableNo).Delete(&entity.Cart{}).Error; err != nil {
		return err
	}

	return dbTx.Commit().Error
}

func (r *OrderingRepo) SubmitCart(detail []entity.Order) error {
	dbTx := r.db.Begin()
	defer dbTx.Rollback()

	for _, v := range detail {
		if err := dbTx.Create(&v).Error; err != nil {
			return err
		}
	}

	return dbTx.Commit().Error
}


	