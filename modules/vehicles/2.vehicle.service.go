package vehicles

import (
	"github.com/pius706975/backend/database/orm/models"
	"github.com/pius706975/backend/helper"
	"gorm.io/gorm"
)

type vehicle_service struct {
	repo Vehicle_Repo
}

func NewVehicleService(repo Vehicle_Repo) *vehicle_service {
	return &vehicle_service{repo}
}

// ADD NEW VEHICLE
func (s *vehicle_service) AddNewVehicle(vehicleData *models.Vehicle) *helper.Response {

	// check : model exists
	modelExists, err := s.repo.ModelExists(vehicleData.Model)
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}
	if modelExists {
		return helper.New("Vehicle model already exists", 400, true)
	}

	data, err := s.repo.AddNewVehicle(vehicleData)
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}

	return helper.New(data, 200, false)
}

// REMOVE VEHICLE
func (s *vehicle_service) RemoveVehicle(vehicleID string) *helper.Response {

	_, err := s.repo.GetVehicleByID(vehicleID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helper.New("Data not found", 404, true)
		} else {
			return helper.New(err.Error(), 500, true)
		}
	}

	err = s.repo.RemoveVehicle(vehicleID)
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}

	result := map[string]string{"Message": "Vehicle has been deleted"}

	return helper.New(result, 200, false)
}

// UPDATE VEHICLE
func (s *vehicle_service) UpdateVehicle(data *models.Vehicle, ID string) *helper.Response {

	var vehicle models.Vehicle

	_, err := s.repo.GetVehicleByID(ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.New("Data not found", 404, true)
		} else {
			helper.New(err.Error(), 500, true)
		}
	}

	// check : model exists
	modelExists, err := s.repo.ModelExists(data.Model)
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}
	if modelExists {
		return helper.New("Vehicle model already exists", 400, true)
	}

	// FIELD VALIDATION
	if data.Model == "" {
		data.Model = vehicle.Model
	}
	if data.Category_ID == "" {
		data.Category_ID = vehicle.Category_ID
	}
	if data.Price == 0 {
		data.Price = vehicle.Price
	}
	if data.Rating == 0 {
		data.Rating = vehicle.Rating
	}
	if data.Stock == 0 {
		data.Stock = vehicle.Stock
	}
	if data.Image == "" {
		data.Image = vehicle.Image
	}
	if data.Location == "" {
		data.Location = vehicle.Location
	}
	if data.Status == "" {
		data.Status = vehicle.Status
	}

	result, err := s.repo.UpdateVehicle(data,ID)
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}

	return helper.New(result, 200, false)

}

// GET ALL VEHICLE
func (s *vehicle_service) GetAllVehicles() *helper.Response {
	
	data, err := s.repo.GetAllVehicles()
	if err != nil {
		return helper.New(err.Error(), 500, true)
	}

	return helper.New(data, 200, false)
}

// GET BY CATEGORY
func (s *vehicle_service) GetByCategory(category string) *helper.Response {

	data, err := s.repo.GetByCategory(category)
	if err != nil {
		return helper.New(err.Error(), 500, true)
	}

	return helper.New(data, 200, false)
}

// GET BY MODEL
func (s *vehicle_service) SearchVehicle(query string) *helper.Response {

	data, err := s.repo.SearchVehicle(query)
	if err != nil {
		return helper.New(err.Error(), 500, true)
	}

	return helper.New(data, 200, false)
}

// GET POPULAR VEHICLE
func (s * vehicle_service) GetPopularVehicle() *helper.Response {

	data, err := s.repo.GetPopularVehicles()
	if err != nil {
		return helper.New(err.Error(), 400, true)
	}

	return helper.New(data, 200, false)
}

func (s *vehicle_service) GetVehicleByID(ID string) *helper.Response {

	data, err := s.repo.GetVehicleByID(ID)
	if err != nil {
		return helper.New(err.Error(), 500, true)
	}

	return helper.New(data, 200, false)
}