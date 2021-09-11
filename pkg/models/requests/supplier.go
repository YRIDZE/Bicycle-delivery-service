package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type SupplierRequest struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (c SupplierRequest) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.Name, validation.Required, validation.Length(1, 30)),
		validation.Field(&c.Image, validation.Required, is.URL),
	)
}
