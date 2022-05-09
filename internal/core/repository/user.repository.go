package repository

import (
	"backend-helpdesk-dores/internal/core/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindByEmail(email string) (*model.User, error)
	FindById(id uuid.UUID) (*model.User, error)
}
