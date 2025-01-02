package middleware

import (
	"github.com/haierkeys/golang-image-upload-service/pkg/app"
	"github.com/haierkeys/golang-image-upload-service/pkg/code"

	"github.com/gin-gonic/gin"
)

func NoFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := app.NewResponse(c)
		response.ToResponse(code.ErrorNotFoundAPI)
		c.Abort()
	}
}
