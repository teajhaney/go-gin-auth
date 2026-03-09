package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)



func  RequiredAdmin() gin.HandlerFunc{
	return func(c *gin.Context){
		role, ok := GetRole(c)
		if !ok{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "UnauthorizedP",
			})
			return
		}
		if !strings.EqualFold(role, "admin"){
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Unauthorized: This action requires admin privileges",
			})
			return
		}
		c.Next()
	}
}
