package model

type OrderDetail struct {
	TableNo uint8 `json:"table_no"`
	FoodName string `json:"food_name"`
	Quantity uint8 `json:"quantity"`
}