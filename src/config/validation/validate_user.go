package validation

import (
	"encoding/json"
	"errors"
	errorhandler "gin-gonic-api/src/config/errorHandler"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *errorhandler.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return errorhandler.NewBadRequestError("Campo com tipo inválido")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []errorhandler.Cause{}

		for _, err := range validation_err.(validator.ValidationErrors) {
			cause := errorhandler.Cause{
				Field:   err.Field(),
				Message: err.Translate(transl),
			}
			errorsCauses = append(errorsCauses, cause)
		}

		return errorhandler.NewBadRequestValidationError("Campos inválidos", errorsCauses)
	} else {
		return errorhandler.NewBadRequestError("Erro de validação")
	}
}

// comente o código deste arquivo explicando cada linha para alguém que não sabe nada de Go, mas entende bem de Typescript e Nestjs
