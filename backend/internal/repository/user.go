package repository

import (
	"backend/internal/models"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) IUserRepository {
	return &repository{db: db}
}

func (r *repository) CreateUser(user *models.User) error {
	_, err := r.db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)

	return err
}

func (r *repository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	return &user, err
}
