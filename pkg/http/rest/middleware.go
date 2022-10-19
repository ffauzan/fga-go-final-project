package rest

import (
	"errors"
	"final-project/pkg/domain"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Gin middleware to validate JWT token
func AuthMiddleware(authService domain.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from header
		token := c.GetHeader("Authorization")
		isBearer := strings.HasPrefix(token, "Bearer ")
		if !isBearer {
			SendErrorResponse(c, errors.New("invalid token format"), http.StatusUnauthorized)
			c.Abort()
			return
		}

		// Remove "Bearer " from token
		token = strings.TrimPrefix(token, "Bearer ")

		// Validate token
		userID, err := authService.ValidateToken(token)
		if err != nil {
			SendErrorResponse(c, err, http.StatusUnauthorized)
			c.Abort()
			return
		}

		// Set userID to context
		c.Set("currentUserID", userID)

		c.Next()
	}
}
