package reservation

import (
	"errors"
	// "fmt"
	"log"
	"time"

	"github.com/pius706975/backend/database/orm/models"
	"github.com/pius706975/backend/libs"
	"gorm.io/gorm"
)

type reserve_repo struct {
	db *gorm.DB
}

func NewReserveRepo(db *gorm.DB) *reserve_repo {
	return &reserve_repo{db}
}

// Create reservation
func (r *reserve_repo) CreateReservation(data *models.Reservation) (*models.Reservation, error) {

	var vehicle models.Vehicle

	err := r.db.Where("vehicle_id = ?", data.VehicleID).First(&vehicle).Error
	if err != nil {
		return nil, err
	}

	// check if the stock is enough for the reservation
	if vehicle.Stock < data.Qt {
		return nil, errors.New("not enough stock")
	}

	// Update stock data
	vehicle.Stock = vehicle.Stock - data.Qt
	err = r.db.Save(&vehicle).Error
	if err != nil {
		return nil, err
	}

	data.PaymentCode = libs.GeneratePaymentCode()

	// Update total payment and the reservation
	data.Total_Payment = (vehicle.Price * data.Qt) * uint(data.ReturnDate.Sub(data.StartDate).Hours()/24)
	err = r.db.Create(data).Error
	if err != nil {
		return nil, err
	}

	// fmt.Println(data.ReservationID)

	// create the reservation history
	history := &models.History{
		ReservationsID: data.ReservationID,
		Status:        "Payment Pending",
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}

	err_ := r.db.Create(history).Error

	if err_ != nil {
		return nil, err_
	}

	return data, nil
}

// Return vehicle
func (r *reserve_repo) ReturnVehicle(reservationID uint) error {
	var vehicle models.Vehicle
	var reservation models.Reservation
	var history models.History

	err := r.db.Where("reservation_id = ? and is_deleted = false", reservationID).First(&reservation).Error
	if err != nil {
		return err
	}

	err = r.db.Where("vehicle_id = ?", reservation.VehicleID).First(&vehicle).Error
	if err != nil {
		return err
	}

	vehicle.Stock = vehicle.Stock + reservation.Qt
	vehicle.UpdatedAt = time.Now()

	err = r.db.Save(&vehicle).Error
	if err != nil {
		return err
	}

	err = r.db.Where("reservations_id = ?", reservationID).First(&history).Error
	if err != nil {
		return err
	}

	history.Status = "Vehicle has returned"
	history.UpdatedAt = time.Now()

	err = r.db.Save(&history).Error
	if err != nil {
		return err
	}

	err = r.db.Save(&vehicle).Error
	if err != nil {
		return err
	}

	reservation.IsDeleted = true

	err = r.db.Save(&reservation).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *reserve_repo) Payment(reservationID, userID uint) error {
	var reservation models.Reservation
	var history models.History

	err := r.db.Where("reservation_id = ? and users_id = ?", reservationID, userID).First(&reservation).Error
	if err != nil {
		log.Println("in")
		return err
	}

	err = r.db.Where("reservations_id = ?", reservationID).First(&history).Error
	if err != nil {
		return err
	}

	// history.IsDeleted == true

	if history.IsDeleted {
		return errors.New("you do not have any reservations")
	}

	reservation.IsPaid = true
	
	err = r.db.Save(reservation).Error
	if err != nil {
		return err
	}

	history.Status = "Vehicle is rented"
	history.UpdatedAt = time.Now()

	err = r.db.Save(&history).Error
	if err != nil {
		return err
	}

	return nil
}
