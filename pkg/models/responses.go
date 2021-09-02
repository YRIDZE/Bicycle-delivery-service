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

type SuppliersResponse struct {
	Restaurants []Supplier `json:"restaurants"`
}

type MenuResponse struct {
	Menu []Product `json:"menu"`
}
