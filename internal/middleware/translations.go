package middleware

import (
	"github.com/haierkeys/golang-image-upload-service/global"

	"github.com/gin-gonic/gin"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.GetHeader("locale")
		if locale == "" {
			locale = "zh"
		}
		trans, found := global.Ut.GetTranslator(locale)

		if found {
			c.Set("trans", trans)
		} else {
			c.Set("trans", "en")
		}
		c.Next()
	}
}
