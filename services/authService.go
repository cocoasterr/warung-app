package services

import (
	"errors"
	"regexp"

	PGRepository "github.com/cocoasterr/net_http/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/net_http/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	BaseService
}

func NewAuthService(repo PGRepository.Repository) *AuthService {
	return &AuthService{
		BaseService: BaseService{
			Repo:  repo,
			Model: &models.Users{},
		},
	}
}
func (s *AuthService) PayloadRegisterCheck(payload map[string]interface{}) error {
	username := payload["username"].(string)
	email := payload["email"].(string)
	password := payload["password"].(string)
	if len(username) < 6 {
		return errors.New("username Characther must more than 6 characters")
	}
	if len(password) < 8 {
		return errors.New("password Characther must more than 6 characters")
	}
	if len(username) < 6 {
		return errors.New("username Characther must more than 6 characters")
	}

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	matched, _ := regexp.MatchString(pattern, email)
	if !matched {
		return errors.New("not valid Email")
	}
	return nil
}
func (s *AuthService) HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
func (s *AuthService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
