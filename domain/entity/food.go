package entity

import (
	"time"

	"gorm.io/gorm"
)

// Food is a struct that represents food entity
type Food struct {
	FoodId      uint           `gorm:"primaryKey"`
	Name        string         `gorm:"unique;not null"`
	Description string         `gorm:"not null"`
	Price       int            `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeleteAt    gorm.DeletedAt `gorm:"index"`
}

// TableName is a function to change the table name
func (Food) TableName() string {
	return "foods"
}

// BeforeCreate is a function to generate ULID before creating a new record
func (f *Food) BeforeCreate(tx *gorm.DB) (err error) {
	f.FoodId = uint(time.Now().Unix())
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
	return
}

// BeforeUpdate is a function to update the updated_at field
func (f *Food) BeforeUpdate(tx *gorm.DB) (err error) {
	f.UpdatedAt = time.Now()
	return
}
