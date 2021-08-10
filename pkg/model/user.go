package model

type User struct {
	ID        int32  `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Delete    string `json:"delete"`
}
