package users

import (
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

	// check : mobile number exists
	mbExists, err := s.repo.MobileNumberExist(userReg.MobileNumber)
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}
	if mbExists {
		return helper.New("Mobile number already exists", 400, true)
	}

	// hashing password
	hashPass, err := libs.HashPassword(userReg.Password)
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}
	
	userReg.Password = hashPass

	data, err := s.repo.Register(userReg)
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}

	return helper.New(data, 200, false)
}

// UPDATE
func (s *user_service) UpdateUser(userData *models.User, id uint) *helper.Response {
	
	var user models.User

	// get user data by id and check if there's and error during data retrieval
	err := s.repo.db.Where("user_id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helper.New("Data not found", 404, true)
		} else {
			return helper.New(err.Error(), 500, true)
		}
	}

	// update password
	if userData.Password != "" {
		// hash new password
		hashPassword, err := libs.HashPassword(userData.Password)
		// check if there's an error during password hashing
		if err != nil {
			return helper.New("Password update failed", 400, true)
		}
		// set the hashed password as new password
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

	// check : mobile number exists
	mbExists, err := s.repo.MobileNumberExist(userData.MobileNumber)
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}
	if mbExists {
		return helper.New("Mobile number already exists", 400, true)
	}
	
	// FIELD VALIDATION
	if userData.FirstName == "" {
		userData.FirstName = user.FirstName
	}
	if userData.LastName == "" {
		userData.LastName = user.LastName
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


	result, err := s.repo.UpdateUser(userData, id)
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}

	return helper.New(result, 200, false)
}

// REMOVE USER
func (s *user_service) RemoveUser(userID uint) *helper.Response {
	
	var user models.User

	err := s.repo.db.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helper.New("Data not found", 404, true)
		} else {
			return helper.New(err.Error(), 500, true)
		}
	}

	err = s.repo.RemoveUser(uint(userID))
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}

	result := map[string]string{"Message": "User has been deleted"}

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
