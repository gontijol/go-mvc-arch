package validation

import (
	"bv-api/src/config/rest_err"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate      = validator.New()
	translateUser ut.Translator
)

func InitUserTranslator() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		translateUser, _ = uni.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, translateUser)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var validationErr validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &validationErr) {
		errorCauses := []rest_err.Causes{}

		for _, err := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: err.Translate(translateUser),
				Field:   err.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("Invalid field(s) input", errorCauses)

	} else {
		return rest_err.NewBadRequestError("Error tryng to convert fields")
	}
}
