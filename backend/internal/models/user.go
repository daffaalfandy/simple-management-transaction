package models

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
