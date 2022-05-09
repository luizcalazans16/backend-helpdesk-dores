package database

import (
	"backend-helpdesk-dores/internal/core/model"
	"database/sql"
)

type UserRepository struct {
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{}
}

func (repository UserRepository) search() {

}

func (repository UserRepository) SearchEmail(email string) (model.User, error) {
	var user model.User
	return user, nil
}
