package models

type Product struct {
	ID          int      `json:"id"`
	SupplierID  int32    `json:"supplier_id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Type        string   `json:"type"`
	Ingredients []string `json:"ingredients"`
	Image       string   `json:"image"`
	Deleted     string   `json:"deleted"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}
