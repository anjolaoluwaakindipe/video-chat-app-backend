package dto

type Login struct {
	Email string `json:"email" validate:"required,gt=6,email"`
	Password string `json:"password" validate:"required"`
}
