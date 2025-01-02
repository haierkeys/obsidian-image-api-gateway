package global

import (
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/validator"

	ut "github.com/go-playground/universal-translator"
)

var (
	Validator *validator.CustomValidator
	Ut        *ut.UniversalTranslator
)
