package models

type Cart struct {
	ID      int32     `json:"id"`
	Product []Product `json:"product"`
}
