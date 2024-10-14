package validator

import (
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ValidateUUID(id string) error {
	return Validate.Var(id, "uuid4") // validate id
}

func ValidateEmail(email string) error {
	return Validate.Var(email, "email") // validate email
}
