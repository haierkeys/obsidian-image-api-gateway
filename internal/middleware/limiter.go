package middleware

import (
	"github.com/haierkeys/custom-image-gateway/pkg/app"
	"github.com/haierkeys/custom-image-gateway/pkg/code"
	"github.com/haierkeys/custom-image-gateway/pkg/limiter"

	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.Face) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToResponse(code.ErrorTooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
