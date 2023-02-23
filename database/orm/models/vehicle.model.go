package models

import "time"

type Vehicle struct {
	VehicleId    uint      `gorm:"primaryKey" json:"id,omitempty" valid:"-"`
	Model        string    `gorm:"type: varchar; not null" json:"model" valid:"type(string),required~Model cannot be empty"`
	CategoriesID uint      `gorm:"not null" json:"category_id" valid:"type(uint)"`
	Category     Category  `gorm:"foreignKey:CategoriesID; reference:CategoryID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" valid:"-"`
	Price        uint      `gorm:"not null" json:"price" valid:"type(uint)"`
	Rating       uint      `json:"rating" valid:"-"`
	Stock        uint      `gorm:"not null" json:"stock" valid:"type(uint)"`
	Image        string    `gorm:"type:varchar" json:"image" valid:"-"`
	Location     string    `gorm:"type:varchar; not null" json:"location" valid:"type(string),required~Location is needed"`
	Status       string    `gorm:"type:varchar; default: Available" json:"status" valid:"-"`
	CreatedAt    time.Time `json:"created_at" valid:"-"`
	UpdatedAt    time.Time `json:"updated_at" valid:"-"`
}

type Vehicles []Vehicle

func (Vehicle) TableName() string {
	return "vehicles"
}
