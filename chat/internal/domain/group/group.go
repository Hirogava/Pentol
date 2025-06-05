package group

type Group struct {
	Id int `json:"id"`
	OwnerId int `json:"owner_id"`
}

type GroupDesc struct {
	GroupData Group `json:"group_data"`
	Name string `json:"name"`
	Description string `json:"description"`
	CreatedAt string `json:"created_at"`
}