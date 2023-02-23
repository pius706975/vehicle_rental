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
	result := r.db.Create(data)

	if result.Error != nil {
		return nil, result.Error
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
func (r *Vehicle_Repo) RemoveVehicle(id uint) error {

	result := r.db.Delete(&models.Vehicle{}, "vehicle_id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// update vehicle
func (r *Vehicle_Repo) UpdateVehicle(data *models.Vehicle, id uint) (*models.Vehicle, error) {

	result := r.db.Model(&data).Where("vehicle_id = ?", id).Updates(&data).Find(&data)
	if result.Error != nil {
		return nil, errors.New("update failed")
	}

	return data, nil
}


// get all
func (r *Vehicle_Repo) GetAllVehicles() (*models.Vehicles, error) {

	var data models.Vehicles

	result := r.db.Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

// get popular
func (r *Vehicle_Repo) GetPopularVehicles() (*models.Vehicles, error) {
	
	var data models.Vehicles

	result := r.db.Where("rating >= ?", 4).Order("rating desc").Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

// get by id
func (r *Vehicle_Repo) GetVehicleByID(id uint) (*models.Vehicle, error) {
	var vehicle models.Vehicle

	result := r.db.First(&vehicle, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &vehicle, nil
}

// get by category
func (r *Vehicle_Repo) GetVehicleByCategory(category string) (*models.Vehicles, error) {
	var data models.Vehicles

	result := r.db.Table("vehicles").Select("vehicles.*, categories.category_name").Joins("inner join categories on vehicles.categories_id = categories.category_id").Where("categories.category_name ILIKE ?", "%"+strings.ToLower(category)+"%").Find(&data)
 
	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

// get by model
func (r *Vehicle_Repo) GetVehicleByModel(model string) (*models.Vehicles, error) {
	var data models.Vehicles

	result := r.db.Where("model ILIKE ?", "%"+strings.ToLower(model)+"%").Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}