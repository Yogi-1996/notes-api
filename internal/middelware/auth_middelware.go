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
			"message": "Authorization is missing/empty",
		})
		return
	}

	tokenString := strings.Trim(authHeader, `"`)

	token, err := jwt.VerifyToken(tokenString)

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid or expired token",
		})
		return
	}

	c.Next()

}
