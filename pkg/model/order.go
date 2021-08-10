package model

type Order struct {
	ID      int32     `json:"id"`
	Address string    `json:"address"`
	Items   []Product `json:"items"`
}
