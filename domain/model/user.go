package model

import (
	"foodOrder/domain/entity"
	_ "time"
)

type RegisterUser struct {
	Username string          `json:"username"`
	Password string          `json:"password"`
	Type     entity.Usertype `json:"user_type"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDetail struct {
	ID        string          `json:"id"`
	Username  string          `json:"username"`
	Type      entity.Usertype `json:"user_type"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
}
