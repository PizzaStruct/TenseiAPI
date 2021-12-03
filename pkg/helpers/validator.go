package helpers

import "github.com/go-playground/validator/v10"

func ValidateStruct(i interface{}) error {
	validate := validator.New()
	err := validate.Struct(i)
	if err != nil {
		return err
	}
	return nil
}
