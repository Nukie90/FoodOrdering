package model

import "github.com/oklog/ulid/v2"

type TableOrder struct {
	TableNo int           `json:"table_no"`
	Detail  []OrderDetail `json:"detail"`
}

type OrderDetail struct {
	OrderId   ulid.ULID `json:"order_id"`
	PreferenceID ulid.ULID `json:"preference_id"`
	FoodName string    `json:"food_name"`
	Quantity uint8     `json:"quantity"`
	Status   string    `json:"status"`
}

type SendRobotRequest struct {
	OrderID ulid.ULID `json:"order_id"`
}

type UpdateOrder struct {
	TableNo uint8 `json:"table_no"`
	OrderID ulid.ULID `json:"order_id"`
	Status string `json:"status"`
	Quantity uint8 `json:"quantity"`
}