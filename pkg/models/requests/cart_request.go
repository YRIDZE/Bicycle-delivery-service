package requests

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	validation "github.com/go-ozzo/ozzo-validation"
)

type CartRequest struct {
	ID     int   `json:"id"`
	UserID int32 `json:"user_id"`
}

type CartProductRequest struct {
	ID       int                   `json:"id"`
	UserID   int32                 `json:"user_id"`
	Products []models.CartProducts `json:"products"`
}

func (c CartProductRequest) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.Products, validation.Required),
	)
}
