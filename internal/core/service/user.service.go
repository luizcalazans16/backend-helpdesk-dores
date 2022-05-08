package service

import (
	"backend-helpdesk-dores/internal/core/model"
	"context"

	"github.com/google/uuid"
)

type UserService struct {
	userRepository UserRepository
}

func (u *UserService) FindUserById(ctx context.Context, id uuid.UUID) *model.User {

}
