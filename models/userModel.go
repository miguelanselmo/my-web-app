package models

type UserModel struct {
	Email    string `validate:"required email"`
	Password string `validate:"required password"`
}
