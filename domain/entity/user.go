package entity

import (
	"errors"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Usertype string

const (
	Cooker Usertype = "cooker"
	Staff Usertype = "staff"
)

// User is a struct that represents user entity
type User struct {
	ID ulid.ULID `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Type Usertype `gorm:"not null oneOf=cooker staff"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
}

// TableName is a function to change the table name
func (User) TableName() string {
	return "users"
}

// BeforeCreate is a function to generate ULID before creating a new record
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	ulid := ulid.MustNew(ulid.Now(), nil)
	u.ID = ulid

	if u.Type != Cooker && u.Type != Staff {
		return errors.New("invalid user type")
    }
	return
}

// BeforeUpdate is a function to update the updated_at field
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}