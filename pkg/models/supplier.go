package models

type Supplier struct {
	ID        int32        `json:"id"`
	Name      string       `json:"name"`
	Type      string       `json:"type"`
	Menu      []Product    `json:"menu"`
	Image     string       `json:"image"`
	WorkHours WorkingHours `json:"workingHours"`
	Deleted   string       `json:"deleted"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
}

type WorkingHours struct {
	Opening string `json:"opening"`
	Closing string `json:"closing"`
}

type SupplierTypes string
