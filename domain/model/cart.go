package model

type AddToCart struct {
	FoodName string `json:"food_name"`
	Quantity uint8    `json:"quantity"`
	UserOrder string `json:"user_order"`
	TableNo uint8 `json:"table_no"`
}

type CartDetail struct {
	TableNo uint8 `json:"table_no"`
	FoodName string `json:"food_name"`
	Quantity uint8 `json:"quantity"`
}