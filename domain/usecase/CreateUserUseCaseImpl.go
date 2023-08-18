package usecase

import (
	"context"
	"fmt"
	"maineBot/boundary/model"
	"maineBot/boundary/repository"
	"maineBot/domain/entity"
	"maineBot/domain/entity/user"
	"time"
)

type CreateUserUseCaseImpl struct {
	repository repository.UserRepository
}

func (c CreateUserUseCaseImpl) Execute(info model.CreateUserInfo) (*user.ID, error) {
	name, nameErr := entity.NameFrom(info.Name)
	if nameErr != nil {
		return nil, nameErr
	}

	age, ageErr := user.AgeFrom(info.Age)
	if ageErr != nil {
		return nil, ageErr
	}

	gender, genderOk := user.GenderFrom(info.Gender)
	if !genderOk {
		return nil, fmt.Errorf("valueSyr isn`t Gender type. Wrong Gender: %q ", genderOk)
	}

	createdAt := time.Now()
	updatedAt := time.Now()

	userStatus, userStatusOk := user.StatusFrom(info.Status)
	if !userStatusOk {
		return nil, fmt.Errorf("valueSyr isn`t UserStatus type. Wrong UserStatus %q", userStatusOk)
	}

	var newUser, newUserErr = user.NewBuilder().
		UserID(*user.NewUserID()).
		Name(name).
		TelegramName(info.TelegramName).
		Age(age).
		Gender(gender).
		CreatedAt(createdAt).
		UpdatedAt(updatedAt).
		UserStatus(userStatus).
		Build()

	if newUserErr != nil {
		return nil, newUserErr
	}

	userID, err := c.repository.InsertUser(context.TODO(), newUser)
	if err != nil {
		return nil, err
	}

	return &userID, nil
}
