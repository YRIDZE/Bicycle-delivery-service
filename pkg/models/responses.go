package models

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserResponse struct {
	ID        int32  `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type TokenResponse struct {
	RefreshToken string `json:"refresh_token"`
}

type CachedTokens struct {
	AccessUID  string `json:"access"`
	RefreshUID string `json:"refresh"`
}

type OrderResponse struct {
	ID       int32           `json:"id"`
	UserID   int32           `json:"user_id"`
	Address  string          `json:"address"`
	Status   string          `json:"status"`
	Products []OrderProducts `json:"products"`
}

type SuppliersResponse struct {
	Suppliers []Supplier `json:"restaurants"`
}

type SupplierResponse struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
	Deleted string `json:"deleted"`
}

type ProductsResponse struct {
	Products []Product `json:"menu"`
}

type ProductResponse struct {
	ID          int      `json:"id"`
	SupplierID  int32    `json:"supplier_id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Type        string   `json:"type"`
	Ingredients []string `json:"ingredients"`
	Image       string   `json:"image"`
}

type CartResponse struct {
	ID     int   `json:"id"`
	UserID int32 `json:"user_id"`
}

type CartProductResponse struct {
	ID       int            `json:"id"`
	UserID   int32          `json:"user_id"`
	Products []CartProducts `json:"products"`
}
