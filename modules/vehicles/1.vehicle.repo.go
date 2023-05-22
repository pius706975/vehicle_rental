package vehicles

import (
	"errors"
	"strings"

	"github.com/pius706975/backend/database/orm/models"
	"gorm.io/gorm"
)

type Vehicle_Repo struct {
	db *gorm.DB
}

func NewVehicleRepo(db *gorm.DB) Vehicle_Repo {
	return Vehicle_Repo{db}
}

// add vehicle
func (r *Vehicle_Repo) AddNewVehicle(data *models.Vehicle) (*models.Vehicle, error) {

	result := r.db.Preload("Category").Create(data).Find(&data).Error
	if result != nil {
		return nil, errors.New("create data failed")
	}

	return data, nil
}

// vehicle model exists
func (r *Vehicle_Repo) ModelExists(model string) (bool, error) {
	var count int64

	err := r.db.Model(&models.Vehicle{}).Where("model = ?", model).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// remove vehicle
func (r *Vehicle_Repo) RemoveVehicle(ID string) error {

	var data models.Vehicle

	result := r.db.Delete(data, "vehicle_id = ?", ID).Error
	if result != nil {
		return result
	}

	return nil
}

// update vehicle
func (r *Vehicle_Repo) UpdateVehicle(data *models.Vehicle, ID string) (*models.Vehicle, error) {

	result := r.db.Model(&data).Preload("Category").Where("vehicle_id = ?", ID).Updates(&data).Find(&data).Error
	if result != nil {
		return nil, errors.New("update failed")
	}

	return data, nil
}

// get all
func (r *Vehicle_Repo) GetAllVehicles() (*models.Vehicles, error) {

	var data models.Vehicles

	result := r.db.Preload("Category").Find(&data).Error
	if result != nil {
		return nil, errors.New("get data failed")
	}

	if len(data) <= 0 {
		return nil, errors.New("vehicle data is empty")
	}

	return &data, nil
}

// get popular
func (r *Vehicle_Repo) GetPopularVehicles() (*models.Vehicles, error) {

	var data models.Vehicles

	result := r.db.Preload("Category").Where("rating >= ?", 4).Order("rating desc").Find(&data).Error

	if result != nil {
		return nil, result
	}

	return &data, nil
}

// get by id
func (r *Vehicle_Repo) GetVehicleByID(ID string) (*models.Vehicle, error) {

	var vehicle models.Vehicle

	result := r.db.Preload("Category").First(&vehicle, "vehicle_id", ID).Error
	if result != nil {
		return nil, result
	}

	return &vehicle, nil
}

func (r *Vehicle_Repo) GetByCategory(category string) (*models.Vehicles, error) {
	
	var data models.Vehicles

	result := r.db.Preload("Category").Where("lower(categories.category_name) ILIKE ?", "%"+strings.ToLower(category)+"%").Joins("inner join categories on vehicles.category_id = categories.category_id").Find(&data).Error
	if result != nil {
		return nil, result
	}

	return &data, nil
}

// get by model
func (r *Vehicle_Repo) SearchVehicle(query string) (*models.Vehicles, error) {

	var data models.Vehicles

	result := r.db.Preload("Category").Where("lower(vehicles.model) ILIKE ? OR lower(categories.category_name) ILIKE ? OR lower(vehicles.location) ILIKE ?", "%"+strings.ToLower(query)+"%", "%"+strings.ToLower(query)+"%", "%"+strings.ToLower(query)+"%").Joins("inner join categories on vehicles.category_id = categories.category_id").Find(&data).Error
	if result != nil {
		return nil, result
	}

	if len(data) <= 0 {
		return nil, errors.New("vehicle not found")
	}

	return &data, nil
}