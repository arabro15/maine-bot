package usecase

import (
	"context"
	"maineBot/boundary/repository"
	"maineBot/domain/entity/user"
)

type ReadUserUseCaseImpl struct {
	repository repository.UserRepository
}

func (r ReadUserUseCaseImpl) FindByID(userIDStr string) (user.User, error) {
	id, err := user.IDFrom(userIDStr)
	if err != nil {
		return user.User{}, err
	}

	userByID, err := r.repository.FindUserByID(context.TODO(), id)
	if err != nil {
		return user.User{}, err
	}

	return userByID, nil
}

func (r ReadUserUseCaseImpl) FindByTelegramName(telegramName string) (user.User, error) {
	userByTelegramName, err := r.repository.FindUserByTelegramName(context.TODO(), telegramName)
	if err != nil {
		return user.User{}, err
	}

	return userByTelegramName, nil
}

func (r ReadUserUseCaseImpl) FindALl() ([]user.User, error) {
	allUsers, err := r.repository.FindAllUser(context.TODO())
	if err != nil {
		return nil, err
	}

	return allUsers, nil
}
