package entity

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

// Food is a struct that represents food entity
type Order struct {
	OrderId   ulid.ULID           `gorm:"primaryKey"`
	Status    string         `gorm:"not null"`
	TableNo   uint8          `gorm:"foreignKey:TableNo"`
	FoodId    uint           `gorm:"foreignKey:FoodId"`
	Quantity  uint8          `gorm:"not null"`
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
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	ulid:= ulid.MustNew(ulid.Now(), entropy)
	o.OrderId = ulid
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
	
	if o.Status == "" {
		o.Status = "cooking"
	}

	return
}

// BeforeUpdate is a function to update the updated_at field
func (o *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	o.UpdatedAt = time.Now()
	return
}
