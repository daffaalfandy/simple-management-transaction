package service

import (
	"backend/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

type IAuth interface {
	Register(user *models.User) (*models.User, error)
	Login(credentials *models.UserLoginParams) (string, error)
	ValidateJWT(tokenString string) (jwt.MapClaims, error)
}
