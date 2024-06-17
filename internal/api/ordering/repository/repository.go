package repository

import (
	"errors"
	"foodOrder/domain/entity"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type OrderingRepo struct {
	db *gorm.DB
}

func NewOrderingRepo(db *gorm.DB) *OrderingRepo {
	return &OrderingRepo{db: db}
}

func (r *OrderingRepo) GetTableNo(tableID ulid.ULID) (uint8, error) {
	dbTx := r.db.Begin()
	defer dbTx.Rollback()

	var table entity.TablePreference
	if err := dbTx.Where("preference_id = ?", tableID).First(&table).Error; err != nil {
		return 0, err
	}

	return table.TableNo, nil
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

	if len(cart) == 0 {
		return nil, errors.New("cart is empty")
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

func (r *OrderingRepo)SubmitCart(detail []entity.Order) error {
	dbTx := r.db.Begin()
	defer dbTx.Rollback()

	for _, v := range detail {
		if err := dbTx.Create(&v).Error; err != nil {
			return err
		}
	}

	return dbTx.Commit().Error
}

func (r *OrderingRepo) GetOrder(tableNo uint8) ([]entity.Order, error) {
	var order []entity.Order
	if err := r.db.Where("table_no = ?", tableNo).Find(&order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderingRepo) GetOrderByID(orderId ulid.ULID) ([]entity.Order, error) {
	var order []entity.Order
	if err := r.db.Where("order_id = ?", orderId).Find(&order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderingRepo) UpdateOrder(order entity.Order) error {
	dbTx := r.db.Begin()
	defer dbTx.Rollback()

	if err := dbTx.Model(&entity.Order{}).Where("order_id = ?", order.OrderId).Update("status", order.Status).Error; err != nil {
		return err
	}

	return dbTx.Commit().Error
}

func (r *OrderingRepo) TableAmount() (uint8, error) {
	var totalTable uint8
	if err := r.db.Table("tables").Select("COUNT(table_no)").Find(&totalTable).Error; err != nil {
		return 0, err
	}

	return totalTable, nil
}