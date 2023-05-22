package users

import (
	"log"
	"os"

	"github.com/pius706975/backend/database/orm/models"
	"github.com/pius706975/backend/helper"
	"github.com/pius706975/backend/libs"
	"gorm.io/gorm"
)

type user_service struct {
	repo User_Repo
}

func NewUserService(repo User_Repo) *user_service {
	return &user_service{repo}
}

// REGISTER
func (s *user_service) Register(userReg *models.User) *helper.Response {

	// check : email exists
	emailExists, err := s.repo.EmailExist(userReg.Email)
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}
	if emailExists {
		return helper.New("Email already exists", 400, true)
	}

	// check : username exists
	usernameExists, err := s.repo.UsernameExists(userReg.Username)
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}
	if usernameExists {
		return helper.New("Username already used", 400, true)
	}

	// hashing password
	hashPassword, err := libs.HashPassword(userReg.Password)
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}

	userReg.Password = hashPassword

	// token verify
	tokenVerify, err := libs.CodeCrypt(32)
	if err != nil {
		return helper.New(err.Error(), 500, true)
	}

	userReg.TokenVerify = tokenVerify

	// send email
	emailData := libs.EmailData{
		URL: os.Getenv("BASE_URL") + "/auth/confirm_email/" + tokenVerify,
		Username: userReg.Username,
		Subject: "Your verification code",
	}

	err = libs.SendEmail(userReg, &emailData)
	if err != nil {
		return helper.New(err.Error(), 500, true)
	}

	log.Println(err)

	data, err := s.repo.Register(userReg)
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}

	result, _ := s.repo.GetByID(data.UserID)
	
	// log.Println(data)

	return helper.New(result, 200, false)
}

// UPDATE
func (s *user_service) UpdateUser(userData *models.User, ID string) *helper.Response {

	var user models.User

	// get user data by id and check if there's and error during data retrieval
	err := s.repo.db.Where("user_id = ?", ID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helper.New("Data not found", 404, true)
		} else {
			return helper.New(err.Error(), 500, true)
		}
	}

	// update password
	if userData.Password != "" {

		hashPassword, err := libs.HashPassword(userData.Password)
		if err != nil {
			return helper.New("Password update failed", 400, true)
		}

		userData.Password = hashPassword
	}

	// check : email exists
	emailExists, err := s.repo.EmailExist(userData.Email)
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}
	if emailExists {
		return helper.New("Email already exists", 400, true)
	}

	// check : username exists
	usernameExists, err := s.repo.UsernameExists(userData.Username)
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}
	if usernameExists {
		return helper.New("Username already used", 400, true)
	}

	// FIELD VALIDATION
	if userData.Name == "" {
		userData.Name = user.Name
	}
	if userData.Username == "" {
		userData.Username = user.Username
	}
	if userData.Email == "" {
		userData.Email = user.Email
	}
	if userData.Password == "" {
		userData.Password = user.Password
	}
	if userData.Gender == "" {
		userData.Gender = user.Gender
	}
	if userData.Address == "" {
		userData.Address = user.Address
	}
	if userData.MobileNumber == "" {
		userData.MobileNumber = user.MobileNumber
	}
	if userData.Role == "" {
		userData.Role = user.Role
	}
	if userData.Image == "" {
		userData.Image = user.Image
	}

	result, err := s.repo.UpdateUser(userData, ID)
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}

	return helper.New(result, 200, false)
}

// REMOVE USER
func (s *user_service) RemoveUser(userID string) *helper.Response {

	_, err := s.repo.GetByID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helper.New("Data not found", 404, true)
		} else {
			return helper.New(err.Error(), 500, true)
		}
	}

	err = s.repo.RemoveUser(userID)
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}

	result := map[string]string{"Message": "Account has been deleted"}

	return helper.New(result, 200, false)
}

// GET ALL USERS
func (s *user_service) GetAllUsers() *helper.Response {

	data, err := s.repo.GetAllUsers()
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}

	return helper.New(data, 200, false)
}

func (s *user_service) GetByID(ID string) *helper.Response {
	
	data, err := s.repo.GetByID(ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helper.New(err.Error(), 404, true)
		} else {
			return helper.New(err.Error(), 500, true)
		}
	}

	return helper.New(data, 200, false)
}