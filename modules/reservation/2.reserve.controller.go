package reservation

import (
	"encoding/json"
	"net/http"

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

	userID, ok := r.Context().Value(middleware.UserID("user")).(string)
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
	err = c.repo.db.Where("vehicle_id = ?", data.Vehicle_ID).First(&vehicle).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data.User_ID = userID

	result, err := c.repo.CreateReservation(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(result)
}

func (c *reserve_ctrl) ReturnVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	reservationID := r.URL.Query().Get("reservation_id")

	err := c.repo.ReturnVehicle(reservationID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte("Vehicle has been returned"))
}

func (c *reserve_ctrl) Payment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	reservationID := r.URL.Query().Get("reservation_id")

	userID, ok := r.Context().Value(middleware.UserID("user")).(string)
	if !ok {
		helper.New("Unauthorized", 401, true).Send(w)
		return
	}

	err_ := c.repo.Payment(reservationID, userID)

	if err_ != nil {
		http.Error(w, err_.Error(), http.StatusBadRequest)
	}

	w.Write([]byte("Vehicle has been paid"))
}

