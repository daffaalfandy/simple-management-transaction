package service

import "backend/internal/models"

type IAuth interface {
	Register(user *models.User) (*models.User, error)
	Login(credentials *models.UserLoginParams) (string, error)
}
