package middleware

import (
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/app"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/code"

	"github.com/gin-gonic/gin"
)

func UserAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		response := app.NewResponse(c)

		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else if s, exist = c.GetQuery("Token"); exist {
			token = s
		} else if s = c.GetHeader("token"); len(s) != 0 {
			token = s
		} else if s = c.GetHeader("Token"); len(s) != 0 {
			token = s
		}

		if token == "" {
			response.ToResponse(code.ErrorNotUserAuthToken)
			c.Abort()
		} else {
			if user, err := app.ParseToken(token); err != nil {
				response.ToResponse(code.ErrorInvalidUserAuthToken)
				c.Abort()
			} else {
				c.Set("user_token", user)
			}
		}

		c.Next()
	}
}
