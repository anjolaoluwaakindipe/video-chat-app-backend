package validation

import (
	"fmt"

	"github.com/go-playground/locales/en"
	 "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"go.uber.org/fx"
)

type ValidationHandler struct {
	validator *validator.Validate
}

func NewValidationHandler() *ValidationHandler {
	
	return &ValidationHandler{validator: validator.New()}
}

func (v *ValidationHandler) ValidateStruct(entity interface{}) map[string]string {
	en := en.New()
	uni := ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")

	var validation *validator.Validate
	validation = validator.New()

	en_translations.RegisterDefaultTranslations(validation, trans)

	

	err := validation.Struct(entity)

	if err != nil {

		errs := make(map[string]string)

		for _, f := range err.(validator.ValidationErrors) {
			err := f.ActualTag()
			if f.Param() != "" {
				err = fmt.Sprintf("%s=%s", err, f.Param())
			}
			fmt.Println(f.Translate(trans))
			errs[f.Field()] = f.Translate(trans)
		}
		return errs
	}

	return nil
}

var Module = fx.Options(
	fx.Provide(NewValidationHandler),
)
