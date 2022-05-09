package service

import (
	"backend-helpdesk-dores/internal/core/model"
	"backend-helpdesk-dores/internal/core/repository"
)

type LoginService struct {
	userRepository repository.UserRepository
}

func (l *LoginService) SearchEmail(email string) (*model.User, error) {
	foundUser, err := l.userRepository.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	return foundUser, nil

}
