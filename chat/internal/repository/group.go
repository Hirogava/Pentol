package repository

import (
	"fmt"

	"github.com/Hirogava/pentol/internal/domain/group"
	"github.com/Hirogava/pentol/internal/domain/message"
	"github.com/Hirogava/pentol/internal/domain/user"
)

func (manager *Manager) CreateGroup(group *group.Group) (error) {
	err := manager.Conn.QueryRow(`INSERT INTO groups (owner_id) VALUES ($1)`, group.OwnerId).Scan(&group.Id)
	return err
}

func (manager *Manager) CreateGroupDesc(group *group.GroupDesc) (error) {
	_, err := manager.Conn.Exec(`INSERT INTO group_desc(group_id, name, description) VALUES ($1, $2, $3)`, group.GroupData.Id, group.Name, group.Description)
	return err
}

func (manager *Manager) DeleteGroup(groupId, ownerId int) (error) {
	var isOwner bool
	err := manager.Conn.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM groups WHERE id = $1 AND owner_id = $2)",
		groupId, ownerId,
	).Scan(&isOwner)

	if err != nil{
    	return err
	}

	if !isOwner{
		return fmt.Errorf("forbidden")
	}

	_, err = manager.Conn.Exec(`DELETE FROM groups WHERE id = $1`, groupId)
 	return err
}

func (manager *Manager) AddUsersToGroup(ownerId int, groupId int, group []*user.Member) (error) {
	var isOwner bool
	err := manager.Conn.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM groups WHERE id = $1 AND owner_id = $2)",
		groupId, ownerId,
	).Scan(&isOwner)

	if err != nil{
		return err
	}

	if !isOwner{
		return fmt.Errorf("forbidden")
	}

	for _, val := range group {
		_, err := manager.Conn.Exec(`INSERT INTO group_users (group_id, user_id) VALUES ($1, $2)`, groupId, val.Id)
		if err != nil {
			return err
		}
	}
   	return nil
}

func (manager *Manager) DeleteUserFromGroup(ownerId, groupId, userId int) (error) {
	var isOwner bool
	err := manager.Conn.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM groups WHERE id = $1 AND owner_id = $2)",
		groupId, ownerId,
	).Scan(&isOwner)

	if err != nil{
    	return err
	}

	if !isOwner{
		return fmt.Errorf("forbidden")
	}

	_, err = manager.Conn.Exec(`DELETE FROM group_users WHERE user_id = $1 AND group_id = $2`, userId, groupId)
  	return err
}

func (manager *Manager) SendMessage(message *message.MessageNew) (error) {
	_, err := manager.Conn.Exec(`INSERT INTO groups_messages (group_id, user_id, created_at, message) VALUES ($1, $2, $3, $4)`, message.ChatID, message.SenderID, message.TS, message.Text)
	return err
}

func (manager *Manager) DeleteMessageFromGroup(groupId, messageId int) (error) {
	_, err := manager.Conn.Exec(`DELETE FROM messages WHERE id = $1 AND group_id = $2`, messageId, groupId)
 	return err
}