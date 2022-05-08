package repository

import (
	"backend-helpdesk-dores/internal/core/model"
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindById(ctx context.Context, id uuid.UUID) (*model.User, error)
}
