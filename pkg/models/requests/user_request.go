package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	ID        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (c UserRequest) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.FirstName, validation.Required, validation.Length(1, 64)),
		validation.Field(&c.LastName, validation.Required, validation.Length(1, 64)),
		validation.Field(&c.Email, validation.Required, is.Email, validation.Length(7, 30)),
		validation.Field(&c.Password, validation.Required, validation.Length(7, 20)),
	)
}

func (c LoginRequest) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.Email, validation.Required, is.Email, validation.Length(7, 30)),
		validation.Field(&c.Password, validation.Required, validation.Length(7, 20)),
	)
}
