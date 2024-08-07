package model

import "github.com/oklog/ulid/v2"

// CreatePayment struct
type CreatePayment struct {
	TableNo  uint     `json:"table"`
}

type Bill struct {
	TableNo  uint     `json:"table"`
	ReceiptID ulid.ULID `json:"receipt_id"`
	PreferenceID  ulid.ULID     `json:"preference_id"`
	Detail  []BillDetail `json:"detail"`
	Total    float64  `json:"total"`
}

type BillDetail struct {
	FoodName string    `json:"food_name"`
	Quantity uint8     `json:"quantity"`
	Price    float64   `json:"price"`
}

type PayBill struct {
	PaymentMethod string `json:"payment_method"`
	Amount 	  float64 `json:"amount"`
}