package repository

import (
	"github.com/Hirogava/pentol/internal/domain/user"

	"golang.org/x/crypto/bcrypt"
)

func (manager *Manager) Login(email string, password string) (user.User, error) {
	var user user.User
	var hashedPassword string

	query := `SELECT id, name, email, password_hash, phone FROM users WHERE email = $1;`
	err := manager.Conn.QueryRow(query, email).Scan(&user.Id, &user.Name, &user.Email, &hashedPassword, &user.PhoneNumber)
	if err != nil {
		return user, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return user, err
	}

	return user, nil
}

func (manager *Manager) Register(user *user.User, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (name, email, password_hash, phone) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	err = manager.Conn.QueryRow(query, user.Name, user.Email, hashedPassword, user.PhoneNumber).Scan(&user.Id)
	if err != nil {
		return err
	}

	return nil
}