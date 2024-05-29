package entity

import (
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

// Cart is a struct that represents order entity
type Cart struct {
	ID        uint      `gorm:"primaryKey"`
	TableNo   uint8     `gorm:"foreignKey:TableNo"`
	UserOrder ulid.ULID `gorm:"foreignKey:UserOrder"`
	FoodId    uint      `gorm:"foreignKey:FoodId"`
	Quantity  uint8     `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeleteAt  gorm.DeletedAt
}

// TableName is a function to change the table name
func (Cart) TableName() string {
	return "carts"
}

// BeforeCreate is a function to generate ULID before creating a new record
func (c *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uint(time.Now().Unix())
	return
}
