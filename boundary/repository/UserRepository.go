package repository

import (
	"context"
	"maineBot/domain/entity/user"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user user.User) (user.ID, error)
	DeleteUserByID(ctx context.Context, userID user.ID) (bool, error)
	UpdateUser(ctx context.Context, user user.User) (bool, error)
	FindUserByID(ctx context.Context, id user.ID) (user.User, error)
	FindUserByTelegramName(ctx context.Context, telegramName string) (user.User, error)
	FindAllUser(ctx context.Context) ([]user.User, error)
}
