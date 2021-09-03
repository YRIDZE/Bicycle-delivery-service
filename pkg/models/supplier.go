package models

type Supplier struct {
	ID      int32     `json:"id"`
	Name    string    `json:"name"`
	Menu    []Product `json:"menu"`
	Image   string    `json:"image"`
	Deleted string    `json:"deleted"`
}
