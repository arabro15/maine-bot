package usecase

import "maineBot/domain/entity/user"

type ReadUserUseCase interface {
	FindByID(userIDStr string) (user.User, error)
	FindByTelegramName(telegramName string) (user.User, error)
	FindALl() ([]user.User, error)
}
