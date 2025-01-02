package app

import (
	"strings"

	"github.com/haierkeys/golang-image-upload-service/global"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// BindAndValid 绑定请求参数并进行验证，支持多语言
func BindAndValid(c *gin.Context, obj interface{}) (bool, ValidErrors) {
	var errs ValidErrors

	// 绑定请求参数到给定的对象
	if err := c.ShouldBindJSON(obj); err != nil {
		return false, errs // 绑定失败，返回 false
	}

	// 使用全局验证器进行验证
	if err := global.Validator.ValidateStruct(obj); err != nil {
		// 如果验证失败，检查错误类型
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			// 获取翻译器
			v := c.Value("trans")
			trans, _ := v.(ut.Translator)

			// 遍历验证错误并进行翻译
			for _, validationErr := range validationErrors {
				translatedMsg := validationErr.Translate(trans) // 翻译错误消息
				errs = append(errs, &ValidError{
					Key:     validationErr.Field(),
					Message: translatedMsg,
				})
			}
		}
		return false, errs // 返回验证错误
	}

	return true, nil // 绑定和验证都成功，返回 true
}
