package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Role not found",
			})
			c.Abort()
			return
		}

		userRole, ok := role.(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid role",
			})
			c.Abort()
			return
		}

		if userRole != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Admin access only",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
