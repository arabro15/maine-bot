package converter

import (
	"errors"
	"maineBot/adapter/repository/model"
	"maineBot/domain/entity"
	"maineBot/domain/entity/user"
	"time"
)

type UserConverter struct{}

func ToModel(user *user.User) model.UserDbModel {
	id := user.Id().Value()
	name := user.Name().Value()
	telegramName := user.TelegramName()
	age := user.Age().Value()
	gender := user.Gender().String()
	createdAt := user.CreatedAt().String()
	updatedAt := user.UpdatedAt().String()
	status := user.Status().String()

	dbModel := model.NewUserDbModel(
		id,
		name,
		telegramName,
		age,
		gender,
		createdAt,
		updatedAt,
		status,
	)

	return *dbModel
}

func ToEntity(userDbModel model.UserDbModel) (*user.User, error) {
	id, err := user.IDFrom(userDbModel.GetId().String())
	if err != nil {
		return nil, errors.New("error is IDFrom userDbModel in UserConverter")
	}

	name, err := entity.NameFrom(userDbModel.GetName())
	if err != nil {
		return nil, errors.New("error is NameFrom userDbModel in UserConverter")
	}

	telegramName := userDbModel.GetTelegramName()
	age, err := user.AgeFrom(userDbModel.GetAge())
	if err != nil {
		return nil, errors.New("error is AgeFrom userDbModel in UserConverter")
	}

	gender, ok := user.GenderFrom(userDbModel.GetGender())
	if ok != true {
		return nil, errors.New("error is GenderFrom userDbModel in UserConverter")
	}

	createdAt, err := time.Parse(userDbModel.GetCreatedAt(), userDbModel.GetCreatedAt())
	if err != nil {
		return nil, errors.New("error is CreatedAt userDbModel in UserConverter")
	}

	updatedAt, err := time.Parse(userDbModel.GetUpdatedAt(), userDbModel.GetUpdatedAt())
	if err != nil {
		return nil, errors.New("error is CreatedAt userDbModel in UserConverter")
	}

	status, ok := user.StatusFrom(userDbModel.GetStatus())
	if ok != true {
		return nil, errors.New("error is UserFrom userDbModel in UserConverter")
	}

	user, err := user.NewBuilder().
		UserID(id).
		Name(name).
		TelegramName(telegramName).
		Age(age).
		Gender(gender).
		CreatedAt(createdAt).
		UpdatedAt(updatedAt).
		UserStatus(status).Build()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
