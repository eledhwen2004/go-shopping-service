package api

type UserUpdatePayload struct {
	ID          string
	Name        string `validate:"omitempty,required"`
	Email       string `validate:"omitempty,required,email"`
	Password    string `validate:"omitempty,required"`
	PhoneNumber string `validate:"omitempty,required"`
}

type UserProfile struct {
	ID    string
	Name  string
	Email string
}
