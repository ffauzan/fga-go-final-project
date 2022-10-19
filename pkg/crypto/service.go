package crypto

import (
	"final-project/pkg/domain"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
}

func NewCryptoService() domain.CryptoService {
	return &service{}
}

func (s *service) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *service) VerifyPassword(plaintext string, hashed string) error {
	log.Println("plaintext: ", plaintext)
	log.Println("hashed: ", hashed)
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plaintext))
}
