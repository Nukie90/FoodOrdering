package entity

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

// Table is a struct that represents payment entity
type Payment struct {
	ReceiptID ulid.ULID `gorm:"primaryKey"`
	PreferenceID ulid.ULID `gorm:"foreignKey:PreferenceID"`
	Total     uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeleteAt  gorm.DeletedAt
}

// TableName is a function to change the table name
func (Payment) TableName() string {
	return "payments"
}

// BeforeCreate is a function to generate ULID before creating a new record
func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	ulid := ulid.MustNew(ulid.Now(), entropy)
	p.ReceiptID = ulid

	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return
}

// BeforeUpdate is a function to update the updated_at field
func (p *Payment) BeforeUpdate(tx *gorm.DB) (err error) {
	p.UpdatedAt = time.Now()
	return
}
