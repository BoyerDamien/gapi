package core

import "github.com/go-playground/validator"

// ValidationErr est le modèle de réponse en cas d'érreur de validation du model
type ValidationErr struct {
	FailedField string
	Tag         string
	Value       string
}

func Validate(model interface{}) []*ValidationErr {
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
