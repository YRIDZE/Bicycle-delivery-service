package models

type LoginResponse struct {
	UserID       int32  `json:"user_id"`
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
	ID               int32           `json:"id"`
	UserID           int32           `json:"user_id"`
	Address          string          `json:"address"`
	PhoneNumber      string          `json:"phone_number"`
	CustomerName     string          `json:"customer_name"`
	CustomerLastname string          `json:"customer_lastname"`
	Status           string          `json:"status"`
	Products         []OrderProducts `json:"products"`
	CreatedAt        string          `json:"created_at"`
}

type SuppliersResponse struct {
	Suppliers []Supplier `json:"suppliers"`
}

type SupplierResponse struct {
	ID        int32        `json:"id"`
	Name      string       `json:"name"`
	Type      string       `json:"type"`
	Image     string       `json:"image"`
	WorkHours WorkingHours `json:"workingHours"`
	Deleted   string       `json:"deleted"`
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
