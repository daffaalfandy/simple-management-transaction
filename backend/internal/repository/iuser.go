package repository

import "backend/internal/models"

type IUserRepository interface {
	CreateUser(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}
