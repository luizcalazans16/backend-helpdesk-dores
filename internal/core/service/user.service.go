package service

import (
	"backend-helpdesk-dores/internal/core/model"
	"backend-helpdesk-dores/internal/core/repository"
	"context"

	"github.com/google/uuid"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (u *UserService) FindUserById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	foundUser, err := u.userRepository.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	return foundUser, nil
}
