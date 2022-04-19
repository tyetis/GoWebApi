package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/tyetis/goapiexample/app/models"
)

func Validator(data interface{}, next func() interface{}) interface{} {
	if err := validator.New().Struct(data); err != nil {
		return &models.BaseResult{Result: false, Message: err.Error()}
	}
	return next()
}
