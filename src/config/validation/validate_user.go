package validation

import (
	// Pacote padrao para erros de parse JSON.
	// Analogia no NestJS: erro quando o payload nao bate com o tipo esperado.
	"encoding/json"
	// Helpers para inspecionar tipos de erro (parecido com `instanceof` em TS).
	"errors"
	// Alias para seu pacote de erro HTTP customizado.
	errorhandler "gin-gonic-api/src/config/errorHandler"

	// Camada de bind/validate do Gin (similar ao ValidationPipe do NestJS).
	"github.com/gin-gonic/gin/binding"
	// Locale ingles para mensagens de validacao.
	"github.com/go-playground/locales/en"
	// Biblioteca de traducao para mensagens do validator.
	ut "github.com/go-playground/universal-translator"
	// Validator usado pelo Gin por baixo dos panos.
	"github.com/go-playground/validator/v10"
	// Mensagens padrao em ingles para o validator.
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	// Instancia global do validator (como um singleton provider no NestJS).
	Validate = validator.New()
	// Tradutor global das mensagens de erro.
	transl ut.Translator
)

// init roda automaticamente quando o pacote e carregado.
// Pense como um bootstrap automatico do modulo.
func init() {
	// Pega o engine de validacao do Gin e confirma o tipo concreto.
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Cria o locale ingles.
		en := en.New()
		// Cria o "hub" de tradutores.
		unt := ut.New(en, en)
		// Seleciona o tradutor "en".
		transl, _ = unt.GetTranslator("en")
		// Registra traducoes padrao (required, min, max, etc).
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

// ValidateUserError converte erros de bind/validacao para seu formato REST padrao.
// Analogia no NestJS: mapear erros para HttpException/BadRequestException.
func ValidateUserError(validation_err error) *errorhandler.RestErr {
	// Erro quando o tipo recebido no JSON nao corresponde ao esperado na struct.
	var jsonErr *json.UnmarshalTypeError
	// Lista de erros de validacao de tags (`binding:"required"`, `min`, etc).
	var jsonValidationError validator.ValidationErrors

	// Se o erro for de tipo invalido de JSON...
	if errors.As(validation_err, &jsonErr) {
		return errorhandler.NewBadRequestError("Campo com tipo invÃ¡lido")
		// Se o erro for de validacao de campos...
	} else if errors.As(validation_err, &jsonValidationError) {
		// Slice (array dinamico) para acumular causas detalhadas.
		errorsCauses := []errorhandler.Cause{}

		// Percorre cada erro de validacao.
		for _, err := range validation_err.(validator.ValidationErrors) {
			// Monta cada causa com nome do campo e mensagem traduzida.
			cause := errorhandler.Cause{
				Field:   err.Field(),
				Message: err.Translate(transl),
			}
			// Adiciona no slice final.
			errorsCauses = append(errorsCauses, cause)
		}

		// Retorna 400 com lista de causas.
		return errorhandler.NewBadRequestValidationError("Campos invÃ¡lidos", errorsCauses)
	} else {
		// Fallback para qualquer outro erro nao previsto.
		return errorhandler.NewBadRequestError("Erro de validaÃ§Ã£o")
	}
}
