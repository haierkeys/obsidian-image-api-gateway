package middleware

import (
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/app"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/code"

	"github.com/gin-gonic/gin"
)

func NoFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := app.NewResponse(c)
		response.ToResponse(code.ErrorNotFoundAPI)
		c.Abort()
	}
}
