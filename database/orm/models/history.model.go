package models

import "time"

type History struct {
	HistoryID      string      `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"history_id,omitempty" valid:"-"`
	
	Reservation_ID string      `gorm:"not null" json:"reservation_id" valid:"uuidv4"`
	Reservation    Reservation `gorm:"foreignKey:Reservation_ID; references:ReservationID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"reservation_data" valid:"-"`
	
	Status         string      `json:"status"`
	Rating         float64     `json:"rating" valid:"range(1|5)"`
	IsDeleted      bool        `gorm:"default: false" json:"IsDeleted"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
}

type Histories []History

func (History) TableName() string {
	return "Histories"
}
