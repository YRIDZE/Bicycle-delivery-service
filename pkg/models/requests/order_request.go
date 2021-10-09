package requests

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	validation "github.com/go-ozzo/ozzo-validation"
)

type OrderRequest struct {
	ID               int32                  `json:"id"`
	UserID           int32                  `json:"user_id"`
	Address          string                 `json:"address"`
	PhoneNumber      string                 `json:"phone_number"`
	CustomerName     string                 `json:"customer_name"`
	CustomerLastname string                 `json:"customer_lastname"`
	PaymentMethod    string                 `json:"payment_method"`
	Status           string                 `json:"status"`
	Products         []models.OrderProducts `json:"products"`
}

func (c OrderRequest) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.Address, validation.Required),
		validation.Field(&c.Products, validation.Required),
		validation.Field(&c.PhoneNumber, validation.Required),
		validation.Field(&c.CustomerName, validation.Required, validation.Length(1, 64)),
		validation.Field(&c.CustomerLastname, validation.Required, validation.Length(1, 64)),
		validation.Field(&c.PaymentMethod, validation.Required),
	)
}
