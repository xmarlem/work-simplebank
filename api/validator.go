package api

import (
	"github.com/go-playground/validator/v10"

	"work-simplebank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is supported
		if util.IsSupportedCurrency(currency) {
			return true
		}
	}

	return false
}
