package vehicles

import (
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
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

	imageName := r.Context().Value("imageName").(string)
	vehicle.Image = imageName

	err := schema.NewDecoder().Decode(&vehicle, r.MultipartForm.Value)
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

	idStr := mux.Vars(r)
	id, ok := idStr["id"]
	if !ok {
		helper.New("Get the ID first", 400, true).Send(w)
		return
	}

	c.svc.RemoveVehicle(id).Send(w)
}

// UPDATE VEHICLE
func (c *vehicle_ctrl) UpdateVehicle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	idStr := mux.Vars(r)
	id, ok := idStr["id"]
	if !ok {
		helper.New("Get the ID first", 400, true).Send(w)
		return
	}

	var data models.Vehicle

	imageName := r.Context().Value("imageName").(string)
	data.Image = imageName

	err := schema.NewDecoder().Decode(&data, r.MultipartForm.Value)
	if err != nil {
		helper.New(err.Error(), 400, true).Send(w)
		return
	}

	c.svc.UpdateVehicle(&data, id).Send(w)
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
func (c *vehicle_ctrl) GetByCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	query := mux.Vars(r)["category"]
	query = strings.ToLower(query)

	c.svc.SearchVehicle(query).Send(w)
}


// // GET BY MODEL
func (c *vehicle_ctrl) SearchVehicle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	query := mux.Vars(r)["query"]
	query = strings.ToLower(query)

	c.svc.SearchVehicle(query).Send(w)
}

func (c *vehicle_ctrl) GetVehicleByID(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-type", "application/json")

	ID := mux.Vars(r)["id"]
	ID = strings.ToLower(ID)

	c.svc.GetVehicleByID(ID).Send(w)
}