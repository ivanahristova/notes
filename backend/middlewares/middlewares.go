package middlewares

import (
	"net/http"
	"notes/backend/utilities/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := token.Valid(c); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "data": "Unauthorized"})
		} else {
			c.Next()
		}
	}
}
