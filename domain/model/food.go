package model

type CreateFood struct {
	Name    string `json:"name"`
	Description string `json:"desc"`
	Price   int    `json:"price"`
}

type FoodDetail struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Price       int    `json:"price"`
}
