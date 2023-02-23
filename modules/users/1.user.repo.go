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
	

	return userData, nil
}

// Update user
func (r *User_Repo) UpdateUser(userData *models.User, userID uint)  (*models.User, error) {

	

	result := r.db.Model(userData).Where("user_id = ?", userID).Updates(&userData).Find(&userData)
	if result.Error != nil {
		return nil, errors.New("update failed")
	}

	userData.Password = ""
	userData.Role = ""

	return userData, nil
}

// Remove user
func (r *User_Repo) RemoveUser(id uint) error {

	result := r.db.Delete(&models.User{}, "user_id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
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

// GET ALL USERS
func (r *User_Repo) GetAllUsers() (*models.Users, error) {

	var user models.Users

	result := r.db.Select("user_id, first_name, last_name, email, gender, address, mobile_number, image").Where("role = ?", "user").Order("created_at DESC").Find(&user)

	if result.Error != nil {
		return nil, errors.New("get data failed")
	}

	return &user, nil
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

// MOBILE NUMBER EXISTS
func (r *User_Repo) MobileNumberExist(mb string) (bool, error) {

	var count int64

	err := r.db.Model(&models.User{}).Where("mobile_number = ?", mb).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
