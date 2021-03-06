package models

type Order struct {
	ID               int32           `json:"id"`
	UserID           int32           `json:"user_id"`
	Address          string          `json:"address"`
	Status           string          `json:"status"`
	PaymentMethod    string          `json:"payment_method"`
	PhoneNumber      string          `json:"phone_number"`
	CustomerName     string          `json:"customer_name"`
	CustomerLastname string          `json:"customer_lastname"`
	Products         []OrderProducts `json:"products"`
	Deleted          string          `json:"deleted"`
	CreatedAt        string          `json:"created_at"`
	UpdatedAt        string          `json:"updated_at"`
}

type OrderProducts struct {
	OrderID   int32   `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Deleted   string  `json:"deleted"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
