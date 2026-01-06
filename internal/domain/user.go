package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

type User struct {
	Id             int
	Name           string
	LastName       string
	Age            int
	Email          string
	Password       string
	RegisteredTime time.Time
}

type SignUpInput struct {
	Name     string `validate:"required,min=2"`
	LastName string `validate:"required,min=3"`
	Age      int    `validate:"required,gte=0,lte=130"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,regexp=^(?=.*[A-Za-z])(?=.*\\d).+$"`
}

func (su *SignUpInput) Validate() error {
	return validate.Struct(su)
}

type SignInInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,regexp=^(?=.*[A-Za-z])(?=.*\\d).+$"`
}

func (si *SignInInput) Validate() error {
	return validate.Struct(si)
}

type UpdateUser struct {
	//ToDo
	// Name     string `validate:"required,min=2"`
	// LastName string `validate:"required,min=3"`
	// Age      int    `validate:"required,gte=0,lte=130"`
	// Email    string `validate:"required,email"`
}
