package models

type Cart struct {
	ID        int            `json:"id"`
	UserID    int32          `json:"user_id"`
	Products  []CartProducts `json:"products"`
	Deleted   string         `json:"deleted"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
}

type CartProducts struct {
	CartID    int     `json:"cart_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
