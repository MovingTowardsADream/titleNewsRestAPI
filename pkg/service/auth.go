package service

import (
	"crypto/sha1"
	"fmt"
	news "titleNewsRestApi"
	"titleNewsRestApi/pkg/repository"
)

const salt = "ie498grhfneo30fnf"

type AuthServices struct {
	repo repository.Authorization
}

func NewAuthServices(repo repository.Authorization) *AuthServices {
	return &AuthServices{repo: repo}
}

func (s *AuthServices) CreateUser(user news.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
