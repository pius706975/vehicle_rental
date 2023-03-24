package category

import (
	"errors"
	"strings"

	"github.com/pius706975/backend/database/orm/models"
	"gorm.io/gorm"
)

type category_repo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) category_repo {
	return category_repo{db}
}

func (r *category_repo) AddCategory(data *models.Category) (*models.Category, error) {

	result := r.db.Create(data).Error

	if result != nil {
		return nil, errors.New("create data faile")
	}

	return data, nil
}

func (r *category_repo) CategoryExists(category string) (bool, error) {

	var data models.Categories
	var count int64

	err := r.db.Model(&models.Category{}).Where("category_name ILIKE ?", "%"+strings.ToLower(category)+"%").Find(&data).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *category_repo) GetAllCategories() (*models.Categories, error) {

	var data models.Categories

	result := r.db.Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	if len(data) == 0 {
		return nil, errors.New("category is empty")
	}

	return &data, nil
}

func (r *category_repo) RemoveCategory(ID string) (*models.Category, error) {

	var data models.Category

	result := r.db.Delete(data, "category_id = ?", ID).Error

	if result != nil {
		return nil, result
	}

	return &data, nil
}
