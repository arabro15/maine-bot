package controller

import (
	"encoding/json"
	"maineBot/adapter/controller/converter"
	"maineBot/adapter/controller/request"
	"maineBot/boundary/usecase"
	"net/http"
)

type UserController struct {
	createUser   usecase.CreateUserUseCase
	readDataUser usecase.ReadUserUseCase
}

func NewUserController(
	createUser usecase.CreateUserUseCase,
	readDataUser usecase.ReadUserUseCase,
) *UserController {
	return &UserController{
		createUser:   createUser,
		readDataUser: readDataUser,
	}
}

func (uc UserController) CreateUser(w http.ResponseWriter, request request.CreateUserRequest) error {
	createUserInfo := converter.CreateUserRequestToModel(request)
	id, err := uc.createUser.Execute(createUserInfo)
	if err != nil {
		w.WriteHeader(400)
		return err
	}
	idBytes, err := json.Marshal(id.Value())
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(idBytes)
	if err != nil {
		return err
	}

	return nil
}
