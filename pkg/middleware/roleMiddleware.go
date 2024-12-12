package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"video-call-project/internal/user"
)

func RoleMiddleware(allRole ...user.ItemRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": "role in token  not found",
			})
			c.Abort()
			return
		}
		roleStr := role.(string)
		// chuyen role sang parse
		roleItem, err := user.ParseStr2ItemRole(roleStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": "role in token  not found",
			})
			c.Abort()
			return
		}

		log.Print(roleItem)
		for _, role := range allRole {
			if role == roleItem {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": "role in token  not found",
		})
		c.Abort()
		return
	}
}
