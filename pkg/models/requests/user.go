package requests

import "net/url"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	ID        int32  `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (a UserRequest) Validate() url.Values {
	errs := url.Values{}

	if a.FirstName == "" {
		errs.Add("firstName", "The firstName field is required!")
	}

	if a.LastName == "" {
		errs.Add("lastName", "The lastName field is required!")
	}

	if a.Email == "" {
		errs.Add("email", "The email field is required!")
	}

	if a.Password == "" {
		errs.Add("password", "The password field is required!")
	}

	if len(a.Password) < 8 || len(a.Password) > 20 {
		errs.Add("password", "The password field must be between 8-20 chars!")
	}

	return errs
}
