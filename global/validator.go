package global

import (
	"github.com/haierkeys/golang-image-upload-service/pkg/validator"

	ut "github.com/go-playground/universal-translator"
)

var (
	Validator *validator.CustomValidator
	Ut        *ut.UniversalTranslator
)
