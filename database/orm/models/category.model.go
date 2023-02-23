package models

import "time"

type Category struct {
	CategoryID   uint      `gorm:"primaryKey" json:"category_id,omitempty" valid:"-"`
	CategoryName string    `json:"category_name" valid:"type(string)"`
	CreatedAt    time.Time `json:"created_at" valid:"-"`
	UpdatedAt    time.Time `json:"updated_at" valid:"-"`
}

type Categories []Category
