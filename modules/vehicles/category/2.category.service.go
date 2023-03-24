package category

import (
	"github.com/pius706975/backend/database/orm/models"
	"github.com/pius706975/backend/helper"
	"gorm.io/gorm"
)

type category_service struct {
	repo category_repo
}

func NewCategoryService(repo category_repo) *category_service {
	return &category_service{repo}
}

// ADD CATEGORY
func (s *category_service) AddCategory(categoryData *models.Category) *helper.Response {

	// check : category exists
	categoryExists, err := s.repo.CategoryExists(categoryData.CategoryName) 
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}
	if categoryExists {
		return helper.New("Category already exists", 400, true)
	}

	data, err := s.repo.AddCategory(categoryData)
	if err != nil {
		return helper.New(err.Error(), 500, true)
	}

	return helper.New(data, 200, false)
}

// GET ALL CATEGORIES
func (s *category_service) GetAllCategories() *helper.Response {

	data, err := s.repo.GetAllCategories()
	if err != nil {
		return helper.New(err.Error(), 500, true)
	}

	return helper.New(data, 200, false)
}

// REMOVE CATEGORY
func (s *category_service) RemoveCategory(ID string) *helper.Response {

	var category models.Category

	err := s.repo.db.Where("category_id = ?", ID).First(&category).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helper.New("Data not found", 404, true)
		} else {
			return helper.New(err.Error(), 500, true)
		}
	}

	_, err = s.repo.RemoveCategory(ID)
	if err != nil {
		return helper.New(err.Error(), 500, true)
	}

	result := map[string]string{"Message": "Category has been deleted"}
	
	return helper.New(result, 200, false)
}