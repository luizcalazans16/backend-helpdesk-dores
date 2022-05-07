package model

type User struct {
	Id       string
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserStruct struct {
	Data []User
}
