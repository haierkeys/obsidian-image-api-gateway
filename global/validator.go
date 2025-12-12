package global

import (
	"github.com/haierkeys/custom-image-gateway/pkg/validator"

	ut "github.com/go-playground/universal-translator"
)

var (
	Validator *validator.CustomValidator
	Ut        *ut.UniversalTranslator
)
