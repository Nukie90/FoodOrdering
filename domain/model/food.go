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
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
