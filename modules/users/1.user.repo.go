package users

import (
	"errors"

	"github.com/pius706975/backend/database/orm/models"
	"gorm.io/gorm"
)

type User_Repo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) User_Repo {
	return User_Repo{db}
}

// Register 
func (r *User_Repo) Register(userData *models.User) (*models.User, error) {

	result := r.db.Create(userData)
	if result.Error != nil {
		return nil, result.Error
	}

	// for hidding fields
	userData.Password = ""
	userData.MobileNumber = ""
	userData.Role = ""
	userData.Image = ""
	

	return userData, nil
}

// Update user
func (r *User_Repo) UpdateUser(userData *models.User, userID string)  (*models.User, error) {

	result := r.db.Model(userData).Where("user_id = ?", userID).Updates(&userData).Find(&userData)
	if result.Error != nil {
		return nil, errors.New("update failed")
	}

	userData.Password = ""
	userData.Role = ""

	return userData, nil
}

func (r *User_Repo) UpdateToken(ID, token string) error {
	
	var data models.User

	err := r.db.Model(data).Where("user_id = ?", ID).Update("token_verify", token).Error
	if err != nil {
		return errors.New("update failed")
	}

	return nil
}

// Remove user
func (r *User_Repo) RemoveUser(ID string) error {

	var data models.User

	result := r.db.Delete(data, "user_id = ?", ID).Error
	if result != nil {
		return result
	}

	return nil
}

// EMAIL EXISTS
func (r *User_Repo) EmailExist(email string) (bool, error) {

	var count int64

	err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// USERNAME EXISTS
func (r *User_Repo) UsernameExists(username string) (bool, error) {
	
	var count int64

	err := r.db.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// get by email
func (r *User_Repo) GetByEmail(email string) (*models.User, error) {

	var data models.User

	result := r.db.First(&data, "email = ?", email)
	if result.Error != nil {
		return nil, errors.New("get data failed")
	}

	return &data, nil
}

// GET BY TOKEN
func (r *User_Repo) GetByToken(token string) (*models.User, error) {
	
	var data models.User
	
	err := r.db.First(&data, "token_verify = ?", token).Error
	if err != nil {
		return nil, errors.New("get data failed")
	}

	return &data, nil
}

// TOKEN EXISTS
func (r *User_Repo) TokenExists(token string) bool {
	
	var data models.User

	err := r.db.First(&data, "token_verify = ?", token).Error
	
	return err == nil
}

// GET ALL USERS
func (r *User_Repo) GetAllUsers() (*models.Users, error) {

	var user models.Users

	result := r.db.Select("user_id, name, username, email, gender, address, mobile_number, image").Where("role = ?", "user").Order("created_at DESC").Find(&user).Error

	if result != nil {
		return nil, errors.New("get data failed")
	}

	return &user, nil
}

func (r *User_Repo) GetByID(ID string) (*models.User, error) {
	
	var data models.User

	result := r.db.Select("user_id, name, username, email, gender, address, mobile_number, image, token_verify, created_at, updated_at").Find(&data, "user_id = ?", ID).Error
	if result != nil {
		return nil, errors.New("get data failed")
	}

	return &data, nil
}