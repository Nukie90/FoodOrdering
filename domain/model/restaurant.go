package model

type CreateRestaurant struct {
	Name        string `json:"name"`
	TableNo    uint8  `json:"table_no"`
}

type AdjustTable struct {
	Name 	  string `json:"name"`
	TableNo    uint8  `json:"table_no"`
}