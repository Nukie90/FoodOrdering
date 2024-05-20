package entity

import (
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
	Usertype Usertype `gorm:"not null"`
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
	ulid := ulid.MustNew(ulid.Timestamp(time.Now()), nil)
	u.ID = ulid
	return
}