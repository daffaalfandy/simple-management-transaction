package service

import (
	"backend/config"
	"backend/internal/models"
	"backend/internal/repository"
	"backend/utils"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	cfg            *config.Config
	userRepository repository.IUserRepository
}

func NewService(
	cfg *config.Config,
	userRepository repository.IUserRepository,
) IAuth {
	return &service{
		cfg:            cfg,
		userRepository: userRepository,
	}
}

func (s *service) Register(user *models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	fmt.Printf("password: %s\n", user.Password)
	user.Password = string(hashedPassword)

	if err := s.userRepository.CreateUser(user); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *service) Login(credentials *models.UserLoginParams) (string, error) {
	user, err := s.userRepository.FindByEmail(credentials.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	credentials.Password = strings.TrimSpace(credentials.Password)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := utils.ValidateJWT(tokenString)
	if err != nil {
		return nil, err
	}

	return token, nil
}
