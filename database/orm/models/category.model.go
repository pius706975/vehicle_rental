package models

import "time"

type Category struct {
	CategoryID   string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"category_id,omitempty" valid:"-"`
	CategoryName string    `json:"category_name" valid:"required~Category name cannot be empty,type(string)"`
	CreatedAt    time.Time `json:"created_at" valid:"-"`
	UpdatedAt    time.Time `json:"updated_at" valid:"-"`
}

type Categories []Category

func (Category) TableName() string {
	return "categories"
}