package models

import "time"

type Vehicle struct {
	VehicleID string `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`
	Model     string `gorm:"type: varchar; not null" json:"model" valid:"type(string),required~Model cannot be empty"`

	Category_ID string   `gorm:"not null" json:"category_id" valid:"uuidv4"`
	Category    Category `gorm:"foreignKey:Category_ID; references:CategoryID; constraint:OnUpdate:Cascade,OnDelete:SET NULL;" json:"category_data" valid:"-"`

	Description string    `json:"description,omitempty" valid:"required~Description is needed"`
	Price       uint      `gorm:"not null" json:"price" valid:"type(uint)"`
	Capacity    uint      `gorm:"not null" json:"capacity" valid:"type(uint)"`
	Status      string    `json:"status,omitempty" valid:"-"`
	Rating      float64   `gorm:"default: 0" json:"rating,omitempty" valid:"-"`
	Stock       uint      `gorm:"not null" json:"stock" valid:"type(uint)"`
	Image       string    `gorm:"type:varchar" json:"image" valid:"-"`
	Location    string    `json:"location,omitempty" valid:"type(string),required~Location is needed"`
	CreatedAt   time.Time `json:"created_at" valid:"-"`
	UpdatedAt   time.Time `json:"updated_at" valid:"-"`
}

type Vehicles []Vehicle

func (Vehicle) TableName() string {
	return "vehicles"
}
