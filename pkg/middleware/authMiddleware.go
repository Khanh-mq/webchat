package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"video-call-project/pkg/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}
		fmt.Println("tokenString:", tokenString)
		//
		//  giai ma token
		token, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized access. Token is either missing or invalid.",
			})
			c.Abort()
			return

		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}
		c.Set("email", claims["email"].(string))
		c.Set("role", claims["role"].(string))
		c.Set("uuid", claims["uuid"])
		c.Next()
	}
}
