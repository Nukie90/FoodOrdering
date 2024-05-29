package entity

import (
	"gorm.io/gorm"

	"time"
)

// Restaurant is a struct that represents restaurant entity
type Restaurant struct {
	RestaurantId uint           `gorm:"primaryKey"`
	Name         string         `gorm:"unique;not null"`
	TotalTable   uint8          `gorm:"not null"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeleteAt     gorm.DeletedAt `gorm:"index"`
}

// TableName is a function to change the table name
func (Restaurant) TableName() string {
	return "restaurants"
}

// BeforeCreate is a function to generate ULID before creating a new record
func (r *Restaurant) BeforeCreate(tx *gorm.DB) (err error) {
	r.RestaurantId = uint(time.Now().Unix())
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
	return
}

// BeforeUpdate is a function to update the updated_at field
func (r *Restaurant) BeforeUpdate(tx *gorm.DB) (err error) {
	r.UpdatedAt = time.Now()
	return
}
