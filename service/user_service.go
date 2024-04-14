package service

import (
	"ai-script-generator/config"
	"ai-script-generator/model"
	"ai-script-generator/views"
	"github.com/charmbracelet/log"
	"github.com/oklog/ulid/v2"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u UserService) CreateUser(req views.UserCreateRequest) (views.UserResponse, error) {
	profile := model.Profile{
		Uid:   ulid.Make().String(),
		Wps:   0,
		Name:  req.Name,
		Email: req.Email,
	}

	account := model.Account{
		Uid:     ulid.Make().String(),
		Profile: profile,
	}

	err := config.DB.Create(&account).Error

	if err != nil {
		log.Error(err)
		return views.UserResponse{}, err
	}

	return views.NewUserResponse(account), nil
}
