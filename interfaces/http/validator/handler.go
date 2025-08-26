package validator

import (
	"go-serviceboilerplate/commons/utils"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(i any) error {
	err := cv.validator.Struct(i)
	if err != nil {
		return utils.NewClientError(err)
	}

	return err
}