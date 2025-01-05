package middleware

import (
	"github.com/haierkeys/obsidian-image-api-gateway/global"

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
			trans, _ := global.Ut.GetTranslator("zh")
			c.Set("trans", trans)
		}
		c.Next()
	}
}
