package entity

import (
	"time"

	"gorm.io/gorm"
)

// Food is a struct that represents food entity
type Table struct {
	TableNo uint8 `gorm:"primaryKey"`
	TableStatus string `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
}

// TableName is a function to change the table name
func (Table) TableName() string {
	return "tables"
}