package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ProductRequest struct {
	ID          int      `json:"id"`
	SupplierID  int32    `json:"supplier_id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Type        string   `json:"type"`
	Ingredients []string `json:"ingredients"`
	Image       string   `json:"image"`
}

func (c ProductRequest) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.Name, validation.Required, validation.Min(2)),
		validation.Field(&c.Price, validation.Required, is.Float),
		validation.Field(&c.Type, validation.Required),
		validation.Field(&c.Ingredients, validation.Required),
		validation.Field(&c.Image, validation.Required, is.URL),
	)
}
