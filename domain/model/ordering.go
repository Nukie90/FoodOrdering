package model

import "github.com/oklog/ulid/v2"

type TableOrder struct {
	TableNo int           `json:"table_no"`
	Detail  []OrderDetail `json:"detail"`
}

type OrderDetail struct {
	OderId   ulid.ULID `json:"order_id"`
	FoodName string    `json:"food_name"`
	Quantity uint8     `json:"quantity"`
	Status   string    `json:"status"`
}

type SendRobotRequest struct {
	OrderID ulid.ULID `json:"order_id"`
}