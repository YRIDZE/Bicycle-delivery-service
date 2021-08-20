package models

type OrderProducts struct {
	OrderID   int32 `json:"order_id"`
	ProductID int32 `json:"product_id"`
	Count     int   `json:"count"`
}

type Order struct {
	ID       int32           `json:"id"`
	UserID   int32           `json:"user_id"`
	Address  string          `json:"address"`
	Status   string          `json:"status"`
	Products []OrderProducts `json:"products"`
}
