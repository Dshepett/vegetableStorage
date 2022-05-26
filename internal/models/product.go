package models

type Product struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  uint   `json:"category_id"`
	Price       uint   `json:"price"`
}
