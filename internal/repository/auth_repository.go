package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lstnprtm/wallet-api/internal/domain"
)

type authRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) domain.AuthRepository {
	return &authRepo{db}
}

func (r *authRepo) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE username = ?", username)
	return &user, err
}

func (r *authRepo) CreateUser(username, hashedPassword string) error {
	_, err := r.db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPassword)
	return err
}
