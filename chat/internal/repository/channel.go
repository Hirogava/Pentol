package repository

import (
	"fmt"

	ch "github.com/Hirogava/pentol/internal/domain/channel"
	"github.com/Hirogava/pentol/internal/domain/message"
)

func (manager *Manager) CreateChannel(channel *ch.Channel) (error) {
	err := manager.Conn.QueryRow(`INSERT INTO channels(owner_id) VALUES ($1) RETURNING id`, channel.OwnerId).Scan(&channel.Id)
	return err
}

func (manager *Manager) CreateChannelDesc(channel *ch.ChannelDesc) (error) {
	_, err := manager.Conn.Exec(`INSERT INTO channel_desc(channel_id, name, description) VALUES ($1, $2, $3)`, channel.ChannelData.Id, channel.Name, channel.Description)
	return err
}

func (manager *Manager) CreatePost(post *message.MessageNew) (error) {
	var isOwner bool
	err := manager.Conn.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM channels WHERE id = $1 AND owner_id = $2)",
		post.ChatID, post.SenderID,
	).Scan(&isOwner)

	if err != nil{
		return err
	}

	if !isOwner{
		return fmt.Errorf("forbidden")
	}

	err = manager.Conn.QueryRow(`INSERT INTO channel_posts(channel_id, message) VALUES ($1, $2) RETURNING id`, post.ChatID, post.Text).Scan(&post.Id)
	return err
}

func (manager *Manager) DeletePost(postId, ownerId int) (error) {
	_, err := manager.Conn.Exec(`DELETE FROM channel_posts WHERE id = $1 AND channel_id IN (SELECT id FROM channels WHERE owner_id = $2)`, postId, ownerId)
	return err
}

func (manager *Manager) AddUserToChannel(userId, channelId, ownerId int) (error) {
	var isOwner bool
	err := manager.Conn.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM channels WHERE id = $1 AND owner_id = $2)",
		channelId, ownerId,
	).Scan(&isOwner)

	if err != nil{
		return err
	}

	if !isOwner{
		return fmt.Errorf("forbidden")
	}

	_, err = manager.Conn.Exec(`INSERT INTO channel_users(user_id, channel_id) VALUES ($1, $2)`, userId, channelId)
	return err
}

func (manager *Manager) DeleteUserFromChannel(userId, channelId, ownerId int) (error) {
	var isOwner bool
	err := manager.Conn.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM channels WHERE id = $1 AND owner_id = $2)",
		channelId, ownerId,
	).Scan(&isOwner)

	if err != nil{
    	return err
	}

	if !isOwner{
		return fmt.Errorf("forbidden")
	}

	_, err = manager.Conn.Exec(`DELETE FROM channel_users WHERE user_id = $1 AND channel_id = $2`, userId, channelId)
 	return err
}