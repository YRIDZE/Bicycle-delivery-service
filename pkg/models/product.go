package models

type ProductType struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	ProductID int32  `json:"product_id"`
}

type Product struct {
	ID          int32    `json:"id"`
	SupplierID  int32    `json:"supplier_id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Type        string   `json:"type"`
	Ingredients []string `json:"ingredients"`
	Logo        string   `json:"logo"`
}
