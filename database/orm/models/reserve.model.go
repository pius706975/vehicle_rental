package models

import "time"

type Reservation struct {
	ReservationID uint      `gorm:"primaryKey" json:"id,omitempty"`
	VehicleID     uint      `gorm:"not null" json:"vehicle_id"`
	Vehicle       Vehicle   `gorm:"foreignKey:VehicleID; reference:VehicleID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UsersID       uint      `gorm:"not null" json:"user_id"`
	User          User      `gorm:"foreignKey:UsersID; reference:UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Qt            uint      `gorm:"not null" json:"qt"`
	StartDate     time.Time `gorm:"type:timestamp; not null" json:"start_date"`
	ReturnDate    time.Time `gorm:"type:timestamp; not null" json:"return_date"`
	Total_Payment uint      `json:"total_payment"`
	PaymentCode   string    `json:"payment_code"`
	IsPaid        bool      `gorm:"default: false" json:"IsPayment"`
	IsDeleted     bool      `gorm:"default: false" json:"IsDeleted"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Reservations []Reservation

func (Reservation) TableName() string {
	return "reservations"
}
