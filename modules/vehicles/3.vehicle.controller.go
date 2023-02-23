package vehicles

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/pius706975/backend/database/orm/models"
	"github.com/pius706975/backend/helper"
)

type vehicle_ctrl struct {
	svc *vehicle_service
}

func NewVehicleCTRL(svc *vehicle_service) *vehicle_ctrl {
	return &vehicle_ctrl{svc}
}

func (c *vehicle_ctrl) AddNewVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var vehicle models.Vehicle

	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		helper.New(err.Error(), 400, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(&vehicle)
	if err != nil {
		helper.New(err.Error(), 400, true).Send(w)
		return
	}

	c.svc.AddNewVehicle(&vehicle).Send(w)
}

// REMOVE VEHICLE
func (c *vehicle_ctrl) RemoveVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		helper.New("Get the ID first", 400, true).Send(w)
		return
	}

	c.svc.RemoveVehicle(uint(id)).Send(w)
}

// UPDATE VEHICLE
func (c *vehicle_ctrl) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		helper.New(err.Error(), 400, true).Send(w)
		return
	}

	var data models.Vehicle
	
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helper.New(err.Error(), 400, true).Send(w)
		return
	}

	c.svc.UpdateVehicle(&data, uint(id)).Send(w)
}

// GET ALL
func (c *vehicle_ctrl) GetAllVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	c.svc.GetAllVehicles().Send(w)
}

// GET POPULAR VEHICLE
func (c *vehicle_ctrl) GetPopularVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	c.svc.GetPopularVehicle().Send(w)
}

// GET BY CATEGORY
func (c *vehicle_ctrl) GetVehicleByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	category := mux.Vars(r)["category"]
	category = strings.ToLower(category)

	c.svc.GetVehicleByCategory(category).Send(w)
}

// // GET BY MODEL
func (c *vehicle_ctrl) GetVehicleByModel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	model := mux.Vars(r)["model"]
	model = strings.ToLower(model)

	c.svc.GetVehicleByModel(model).Send(w)
}