package validation

import (
	"fmt"

	validator "github.com/go-playground/validator/v10"
)

type validation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type validationResponse struct {
	Validations []validation `json:"validations"`
}

func generateValidationMessage(field string, rule string) (message string) {
	switch rule {
	case "required":
		return fmt.Sprintf("Field '%s' is '%s'.", field, rule)
	default:
		return fmt.Sprintf("Field '%s' is not valid.", field)
	}
}

// GenerateValidationResponse ...
func GenerateValidationResponse(err error) (response validationResponse) {
	for _, v := range err.(validator.ValidationErrors) {
		field, rule := v.Field(), v.Tag()
		response.Validations = append(response.Validations, validation{Field: field, Message: generateValidationMessage(field, rule)})
	}
	return
}
