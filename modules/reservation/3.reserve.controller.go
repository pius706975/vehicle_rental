package reservation

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pius706975/backend/database/orm/models"
	"github.com/pius706975/backend/helper"
	"github.com/pius706975/backend/middleware"
)

type reserve_ctrl struct {
	repo *reserve_repo
}

func NewReserveCTRL(repo *reserve_repo) *reserve_ctrl {
	return &reserve_ctrl{repo}
}

// Create reservation
func (c *reserve_ctrl) CreateReservation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	userID, ok := r.Context().Value(middleware.UserID("user")).(uint)
	if !ok {
		helper.New("Unauthorized", 401, true).Send(w)
		return
	}

	var data models.Reservation

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vehicle := models.Vehicle{}
	err = c.repo.db.First(&vehicle, data.VehicleID).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// err = c.repo.db.Save(&vehicle).Error
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	data.UsersID = userID

	result, err := c.repo.CreateReservation(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(result)
}

func (c *reserve_ctrl) ReturnVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	reservationID, err := strconv.Atoi(r.URL.Query().Get("reservation_id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err_ := c.repo.ReturnVehicle(uint(reservationID))

	if err_ != nil {
		http.Error(w, err_.Error(), http.StatusBadRequest)
	}

	w.Write([]byte("Vehicle has been returned"))
}

func (c *reserve_ctrl) Payment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	reservationID, err := strconv.Atoi(r.URL.Query().Get("reservation_id"))

	userID, ok := r.Context().Value(middleware.UserID("user")).(uint)
	if !ok {
		helper.New("Unauthorized", 401, true).Send(w)
		return
	}
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err_ := c.repo.Payment(uint(reservationID), userID)

	if err_ != nil {
		http.Error(w, err_.Error(), http.StatusBadRequest)
	}

	w.Write([]byte("Vehicle has been paid"))
}

