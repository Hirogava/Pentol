package user

type User struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserDesc struct {
	Id int `json:"id"`
	Auth_user_id int `json:"auth_user_id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type Member struct {
	Id int `json:"id"`
	Name string `json:"name"`
}