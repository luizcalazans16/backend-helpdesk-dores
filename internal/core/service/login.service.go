package service

import (
	"backend-helpdesk-dores/internal/core/model"
	"backend-helpdesk-dores/internal/core/repository"
	"context"
)

type LoginService struct {
	userRepository repository.UserRepository
}

func (l *LoginService) SearchEmail(ctx context.Context, email string) (*model.User, error) {
	foundUser, err := l.userRepository.FindByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return foundUser, nil

}
