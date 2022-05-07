package repository

import (
	"backend-helpdesk-dores/src/model"
	"database/sql"
)

type User struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *User {
	return &User{db}
}

func (repository User) search() {

}

func (repository User) SearchEmail(email string) (model.User, error) {
	var user model.User
	return user, nil
}
