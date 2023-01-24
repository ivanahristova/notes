package middlewares

import (
	"net/http"
	"notes/backend/utilities/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := token.Valid(c); err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}
