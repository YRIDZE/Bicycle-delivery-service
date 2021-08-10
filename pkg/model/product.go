package model

type Product struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        int32   `json:"type"`
	Cost        float64 `json:"cost"`
}
