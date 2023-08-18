package auth

import (
	"os"

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

// LOGIN
func (s *auth_service) Login(body *models.User) *helper.Response {

	user, err := s.repo.GetByEmail(body.Email)
	if err != nil {
		return helper.New("Email or password is incorrect", 401, true)
	}

	if libs.CheckPassword(body.Password, user.Password) {
		return helper.New("Email or password is incorrect", 401, true)
	}

	if !user.IsActive {
		return helper.New("Your account is not verified", 401, true)
	}

	jwt := libs.NewToken(user.UserID, user.Role)

	token, err := jwt.CreateToken()
	if err != nil {
		helper.New(err.Error(), 500, true)
	}

	return helper.New(tokenRes{Token: token}, 200, false)

}

// VERIFY EMAIL
func (s *auth_service) VerifyEmail(token string) *helper.Response {

	tokenExists := s.repo.TokenExists(token)
	if !tokenExists {
		return helper.New("Verification failed", 401, true)
	}

	user, err := s.repo.GetByToken(token)
	if err != nil {
		return helper.New("User does not exist", 401, true)
	}

	if user.IsActive {
		return helper.New("The account has been registered", 401, true)
	}

	var data models.User

	data.IsActive = true

	_, err = s.repo.UpdateUser(&data, user.UserID)
	if err != nil {
		return helper.New("User does not exist", 401, true)
	}

	res := map[string]string{"message": "Email has been verified"}

	return helper.New(res, 200, false)
}

// RESEND EMAIL
func (s *auth_service) ResendEmail(data *models.User) *helper.Response {

	emailExists, err := s.repo.EmailExist(data.Email)
	if err != nil {
		return helper.New(err.Error(), 401, true)
	}
	if emailExists {
		return helper.New("Email is not registered", 401, true)
	}

	user, err := s.repo.GetByEmail(data.Email)
	if err != nil {
		return helper.New("User does not exist", 401, true)
	}

	tokenVerify, err := libs.CodeCrypt(32)
	if err != nil {
		return helper.New(err.Error(), 500, true)
	}

	data.TokenVerify = tokenVerify

	emailData := libs.EmailData{
		URL:      os.Getenv("BASE_URL") + "/auth/confirm_email" + tokenVerify,
		Username: data.Username,
		Subject:  "Your verification code",
	}

	err = libs.SendEmail(data, &emailData)
	if err != nil {
		return helper.New(err.Error(), 500, true)
	}

	err = s.repo.UpdateToken(user.UserID, tokenVerify)
	if err != nil {
		return helper.New(err.Error(), 500, true)
	}

	res := map[string]string{"message": "Email verification resent"}

	return helper.New(res, 200, false)
}
