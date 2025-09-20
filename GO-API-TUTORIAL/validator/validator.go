package validator

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type validationError struct {
	Namespace       string `json:"namespace"` // can differ when a custom TagNameFunc that is registered or
	Field           string `json:"field"`     // by passing alt name to ReportError like below
	StructNamespace string `json:"structNamespace"`
	StructField     string `json:"structField"`
	Tag             string `json:"tag"`
	ActualTag       string `json:"actualTag"`
	Kind            string `json:"kind"`
	Type            string `json:"type"`
	Value           string `json:"value"`
	Param           string `json:"param"`
	Message         string `json:"message"`
}

func ValidateStruct[T any](structure T) {

	var validate = validator.New(validator.WithRequiredStructEnabled())

	book := &structure

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(book)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			fmt.Println(err)
			return
		}

		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			for _, err := range validateErrs {
				e := validationError{
					Namespace:       err.Namespace(),
					Field:           err.Field(),
					StructNamespace: err.StructNamespace(),
					StructField:     err.StructField(),
					Tag:             err.Tag(),
					ActualTag:       err.ActualTag(),
					Kind:            fmt.Sprintf("%v", err.Kind()),
					Type:            fmt.Sprintf("%v", err.Type()),
					Value:           fmt.Sprintf("%v", err.Value()),
					Param:           err.Param(),
					Message:         err.Error(),
				}

				indent, err := json.MarshalIndent(e, "", "  ")
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(string(indent))
			}
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}
}
