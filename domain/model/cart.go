package model

type AddToCart struct {
	FoodName string `json:"food_name"`
	Quantity uint8    `json:"quantity"`
	TableID string `json:"table_id"`
}

type CartDetail struct {
	TableNo uint8 `json:"table_no"`
	FoodName string `json:"food_name"`
	Quantity uint8 `json:"quantity"`
}