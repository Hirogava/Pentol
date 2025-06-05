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

func (manager *Manager) GetUsername(id int) (string, error) {
	var username string

	err := manager.Conn.QueryRow(`SELECT name FROM users WHERE id = $1`, id).Scan(&username)
	if err != nil {
		return "", err
	}

	return username, nil
}

func (manager *Manager) GetUser(id int) (*user.UserDesc, error) {
	var user user.UserDesc
	err := manager.Conn.QueryRow(`SELECT users.id, users.name, users_desc.description FROM users LEFT JOIN users_desc ON users.id = users_desc.user_id WHERE users.id = $1`, id).Scan(&user.Id, &user.Name, &user.Description)
	if err != nil {
		return nil, err
	}

	return &user, nil
}