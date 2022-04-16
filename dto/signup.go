package dto

type SignUp struct {
	Email     string `json:"email" validate:"email,required"`
	Password  string `json:"password" validate:"required,min=6"`
	Username  string `json:"username" validate:"required,min=8"`
	Firstname string `json:"firstname" valiedate:"required,min=2"`
	Lastname  string `json:"lastname" valiedate:"required,min=2"`
}
