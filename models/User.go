package models

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type AddUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type DeleteUser struct {
	Id int `json:"id"`
}
