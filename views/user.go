package views

import "ai-script-generator/model"

type UserCreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponse struct {
	Uid   string `json:"uid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserResponse(account model.Account) UserResponse {
	return UserResponse{
		Uid:   account.Uid,
		Name:  account.Profile.Name,
		Email: account.Profile.Email,
	}
}
