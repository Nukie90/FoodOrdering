package entity

import (
	"time"

	"gorm.io/gorm"
)

// Food is a struct that represents food entity
type Order struct {
	Table 	  string    `gorm:"not null"`
	OrderTime time.Time `gorm:"autoCreateTime"`
	OrderList string    `gorm:"not null"`
	Total     int       `gorm:"not null"`
	Finsihed  bool      `gorm:"not null"`
	DeleteOrder gorm.DeletedAt `gorm:"index"`
}

// TableName is a function to change the table name
func (Order) TableName() string {
	return "orders"
}
