package models

import (
	"time"

	"github.com/go-playground/validator"
)

type Model struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func ValidateModel(model interface{}) []*ValidationErr {
	var errors []*ValidationErr
	validate := validator.New()
	err := validate.Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationErr
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
