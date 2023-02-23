package auth

import (
	"github.com/pius706975/backend/database/orm/models"
	"github.com/pius706975/backend/helper"
	"github.com/pius706975/backend/libs"
	"github.com/pius706975/backend/modules/users"
)

type auth_service struct {
	repo users.User_Repo
}

type tokenRes struct {
	Token string `json:"token"`
}

func NewAuthService(repo users.User_Repo) *auth_service {
	return &auth_service{repo}
}

func (s *auth_service) Login(body *models.User) *helper.Response {

	user, err := s.repo.GetByEmail(body.Email)

	if err != nil {
		return helper.New("Email or password is incorrect", 401, true)
	}

	if libs.CheckPassword(body.Password, user.Password) {
		return helper.New("Email or password is incorrect", 401, true)
	}

	jwt := libs.NewToken(user.UserID, user.Role)
	
	token, err := jwt.CreateToken()
	if err != nil {
		helper.New(err.Error(), 500, true)
	}

	return helper.New(tokenRes{Token: token}, 200, false)

}