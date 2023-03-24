package models

import "time"

type Reservation struct {
	ReservationID string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"reservation_id,omitempty" valid:"-"`
	
	// VehicleID     string    `gorm:"foreignKey:VehicleID; references:VehicleID; not null;" json:"vehicle_id" valid:"uuidv4"`
	// Vehicle       Vehicle   `gorm:"foreignKey:VehicleID; references:VehicleID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"vehicle_data" valid:"-"`
	Vehicle_ID     string    `gorm:"not null" json:"vehicle_id" valid:"uuidv4"`
	Vehicle       Vehicle   `gorm:"foreignKey:Vehicle_ID; references:VehicleID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"vehicle_data" valid:"-"`
		
	User_ID       string    `gorm:"not null" json:"user_id" valid:"uuidv4"`
	User          User      `gorm:"foreignKey:User_ID; references:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user_data,omitempty" valid:"-"`
	
	// UsersID       string    `gorm:"foreignKey:UserID; references:UserID;" json:"user_id" valid:"uuidv4"`
	// User          User      `gorm:"foreignKey:UsersID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user_data,omitempty" valid:"-"`

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