package models

type ProductType struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	ProductID int32  `json:"product_id"`
}

type Product struct {
	ID          int32   `json:"id"`
	SupplierID  int32   `json:"supplier_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        int32   `json:"type"`
	Price       float64 `json:"price"`
}
