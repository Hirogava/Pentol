package repository

import (
	"github.com/Hirogava/pentol/internal/domain/user"

	"golang.org/x/crypto/bcrypt"
)

func (manager *Manager) Login(email string, password string) (user.User, error) {
	var user user.User
	var hashedPassword string

	query := `SELECT id, email, password_hash FROM auth_users WHERE email = $1;`
	err := manager.Conn.QueryRow(query, email).Scan(&user.Id, &user.Email, &hashedPassword)
	if err != nil {
		return user, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return user, err
	}

	return user, nil
}

func (manager *Manager) Register(user *user.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO auth_users (email, password_hash) VALUES ($1, $2, $3) RETURNING id;`
	err = manager.Conn.QueryRow(query, user.Email, hashedPassword).Scan(&user.Id)
	if err != nil {
		return err
	}

	return nil
}