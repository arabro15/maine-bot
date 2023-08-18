package converter

import (
	"maineBot/adapter/controller/request"
	"maineBot/boundary/model"
)

type UserRequestConverter struct{}

func CreateUserRequestToModel(requestModel request.CreateUserRequest) model.CreateUserInfo {
	name := requestModel.Name()
	telegramName := requestModel.TelegramName()
	age := requestModel.Age()
	gender := requestModel.Gender()
	createdAt := requestModel.CreatedAt()
	updatedAt := requestModel.UpdatedAt()
	status := requestModel.Status()

	return *model.NewCreateUserInfo(name, telegramName, age, gender, createdAt, updatedAt, status)
}
