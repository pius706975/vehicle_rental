package models

import "time"

type History struct {
	HistoryID      uint        `gorm:"primaryKey" json:"id,omitempty"`
	ReservationsID uint        `json:"reservation_id"`
	Reservation    Reservation `gorm:"foreignKey:ReservationsID; reference:ReservationID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Status         string      `json:"status"`
	IsDeleted      bool        `gorm:"default: false" json:"IsDeleted"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
}

type Histories []History

func (History) TableName() string {
	return "Histories"
}
