package entity

import (
	"gorm.io/gorm"

	"time"
)

// Restaurant is a struct that represents restaurant entity
type Table struct {
	TableNo      uint8           `gorm:"primaryKey"`
	Status 	 string          `gorm:"not null"`
	CreatedAt time.Time     `gorm:"autoCreateTime"`
	UpdatedAt time.Time     `gorm:"autoUpdateTime"`
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}

// TableName is a function to change the table name
func (Table) TableName() string {
	return "tables"
}

// BeforeCreate is a function to generate ULID before creating a new record
func (t *Table) BeforeCreate(tx *gorm.DB) (err error) {
	t.Status = "available"

	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	return
}

// BeforeUpdate is a function to update the updated_at field
func (t *Table) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = time.Now()
	return
}

