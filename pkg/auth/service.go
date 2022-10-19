package auth

import (
	"final-project/pkg/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	UserID uint
	jwt.StandardClaims
}

type service struct {
}

func NewAuthService() domain.AuthService {
	return &service{}
}

// GenerateToken is a function to generate JWT token
func (s *service) GenerateToken(userID uint) (string, error) {
	// Set custom claims
	claims := JwtCustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// ValidateToken is a function to validate JWT token
