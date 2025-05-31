package repository

import (
	"github.com/Hirogava/pentol/internal/domain/user"
)


func (manager *Manager) CreateUser(user *user.UserDesc) error {
	var UserId int

	err := manager.Conn.QueryRow(`INSERT INTO users (name, auth_user_id) VALUES ($1, $2) RETURNING id`, user.Name, user.Auth_user_id).Scan(&UserId)
	if err != nil {
		return err
	}

	_, err = manager.Conn.Exec(`INSERT INTO users_desc (user_id, description) VALUES ($1, $2)`, UserId, user.Description)
	return err
}