package service

import (
	"backend-helpdesk-dores/internal/core/model"
	"backend-helpdesk-dores/internal/core/repository"

	"github.com/google/uuid"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (u *UserService) FindUserById(id uuid.UUID) (*model.User, error) {
	foundUser, err := u.userRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	return foundUser, nil
}
