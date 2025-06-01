package chat

type Chat struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	User1Id int `json:"user1_id"`
	User1Username string `json:"user1_username"`
	User2Id int `json:"user2_id"`
	User2Username string `json:"user2_username"`
}