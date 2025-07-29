package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RBAC(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User role information missing"})
			return
		}

		hasAccess := false
		for _, role := range requiredRoles {
			if role == userRole {
				hasAccess = true
				break
			}
		}

		if !hasAccess {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions."})
			return
		}

		c.Next()
	}
}
