package repository

import (
	"github.com/Hirogava/pentol/internal/domain/message"
	"github.com/Hirogava/pentol/internal/domain/chat"
)

func (manager *Manager) CreateChat(userChat *chat.Chat) error{
	err := manager.Conn.QueryRow(`INSERT INTO chats (user_1_id, user_2_id) VALUES ($1, $2) RETURNING id`, userChat.User1Id, userChat.User2Id).Scan(&userChat.Id)
	if err != nil {
		return err
	}

	_, err = manager.Conn.Exec(`INSERT INTO chat_desc (chat_id, name, description) VALUES ($1, $2, $3)`, userChat.Id, userChat.Name, userChat.Description)
	return err
}

func (manager *Manager) CreateMessage(message *message.MessageNew) error{
	err := manager.Conn.QueryRow(`INSERT INTO messages (chat_id, sender_id, created_at, message) VALUES ($1, $2. $3, $4) RETURNING id`, message.ChatID, message.SenderID, message.TS, message.Text).Scan(&message.Id)
	return err
}

func (manager *Manager) GetUserChats(userID int) ([]chat.Chat, error){
	chats := []chat.Chat{}
	rows, err := manager.Conn.Query(`SELECT c.id, c.user_1_id, c.user_2_id, cd.name, cd.description
		FROM chats AS c
		INNER JOIN chat_desc AS cd ON cd.chat_id = c.id
		WHERE c.user_1_id = $1`, userID)
	
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var chat chat.Chat
		err = rows.Scan(&chat.Id, &chat.User1Id, &chat.User2Id, &chat.Name, &chat.Description)
		if err != nil {
			return chats, err
		}
		chats = append(chats, chat)
	}

	return chats, nil
}

func (manager *Manager) DeleteMessageFromChat(id int) error{
	_, err := manager.Conn.Exec(`DELETE FROM chat_messages WHERE id = $1`, id)
	return err
}

func (manager *Manager) DeleteChat(id int) error{
	_, err := manager.Conn.Exec(`DELETE FROM chats WHERE id = $1`, id)
	if err != nil{
		return err
	}

	_, err = manager.Conn.Exec(`DELETE FROM chat_desc WHERE chat_id = $1`, id)
	if err != nil{
		return err
	}

	_, err = manager.Conn.Exec(`DELETE FROM chat_messages WHERE chat_id = $1`, id)
	return err
}