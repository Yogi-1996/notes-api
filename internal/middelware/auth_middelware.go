package middelware

import (
	"net/http"
	"strings"

	"github.com/Yogi-1996/notes-backend/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AunthMiddelware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Authorization header is missing/empty",
		})
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Authorization header format must be Bearer",
		})
		return
	}

	tokenString := parts[1]

	token, claims, err := jwt.VerifyToken(tokenString)

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid or expired token",
		})
		return
	}

	c.Set("UserID", claims.UserID)

	c.Next()

}
