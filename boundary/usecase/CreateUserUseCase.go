package usecase

import (
	"maineBot/boundary/model"
	"maineBot/domain/entity/user"
)

type CreateUserUseCase interface {
	Execute(info model.CreateUserInfo) (*user.ID, error)
}
