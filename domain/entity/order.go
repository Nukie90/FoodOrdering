package entity

import (
	"time"

	"gorm.io/gorm"
)

// Food is a struct that represents food entity
type Order struct {
	OrderId   uint           `gorm:"primaryKey"`
	Status   string         `gorm:"not null oneOf=cooking done"`
	TableNo Table `gorm:"foreignKey:TableNo"`
	FoodId Food `gorm:"foreignKey:FoodId"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}

// TableName is a function to change the table name
func (Order) TableName() string {
	return "orders"
}

// BeforeCreate is a function to generate ULID before creating a new record
func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.OrderId = uint(time.Now().Unix())
	return
}

// BeforeUpdate is a function to update the updated_at field
func (o *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	o.UpdatedAt = time.Now()
	return
}
