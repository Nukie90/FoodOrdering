package entity

import (
	"math/rand"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"

	"time"
)

// Restaurant is a struct that represents restaurant entity
type TablePreference struct {
	TableNo  uint8           `gorm:"primaryKey"`
	PreferenceID ulid.ULID   `gorm:"primaryKey"`
	Status   string          `gorm:"not null"`
	CreatedAt time.Time     `gorm:"autoCreateTime"`
	UpdatedAt time.Time     `gorm:"autoUpdateTime"`
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}

// TableName is a function to change the table name
func (TablePreference) TableName() string {
	return "table_preferences"
}

// BeforeCreate is a function to generate ULID before creating a new record
func (t *TablePreference) BeforeCreate(tx *gorm.DB) (err error) {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	ulid := ulid.MustNew(ulid.Now(), entropy)
	t.PreferenceID = ulid

	t.Status = "active"

	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	return
}