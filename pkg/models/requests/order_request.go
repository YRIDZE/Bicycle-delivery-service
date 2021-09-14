package requests

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	validation "github.com/go-ozzo/ozzo-validation"
)

type OrderRequest struct {
	ID       int32                  `json:"id"`
	UserID   int32                  `json:"user_id"`
	Address  string                 `json:"address"`
	Status   string                 `json:"status"`
	Products []models.OrderProducts `json:"products"`
}

func (c OrderRequest) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.Address, validation.Required),
		validation.Field(&c.Products, validation.Required),
	)
}
