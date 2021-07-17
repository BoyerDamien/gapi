package models

import (
	"encoding/json"
)

type ErrResponse struct {

	// Message d'erreur
	// example: error
	// required: true
	Message string `json:"message"`
}

// ValidationErr est le modèle de réponse en cas d'érreur de validation du model
//
// swagger:model
type ValidationErr struct {
	FailedField string
	Tag         string
	Value       string
}

func (s *ValidationErr) ToString() string {
	res, _ := json.Marshal(s)
	return string(res)
}
