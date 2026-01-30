package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// User defines a user struct with validation tags
type User struct {
	Name     string `validate:"required,min=2,max=50"`
	Email    string `validate:"required,email"`
	Age      int    `validate:"min=0,max=150"`
}

func main() {
	validate := validator.New()

    // It is recommended to use the WithRequiredStructEnabled option
    // for behavior that will become default in v11+
    // validate := validator.New(validator.WithRequiredStructEnabled())

	user := User{
		Name:  "A", // Invalid: min length is 2
		Email: "invalid-email", // Invalid: not a valid email format
		Age:   180, // Invalid: max age is 150
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println("Validation errors:")
		// Iterate over validation errors
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Field: %s, Tag: %s, Value: %v\n", err.Field(), err.Tag(), err.Value())
		}
	} else {
		fmt.Println("Struct is valid")
	}
}
