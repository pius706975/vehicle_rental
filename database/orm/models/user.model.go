package models

import (
	"time"
)

type User struct {
	UserID       uint      `gorm:"primarykey;auto_increment" json:"id,omitempty" valid:"-"`
	FirstName    string    `json:"first_name" valid:"type(string),required~First name cannot be empty"`
	LastName     string    `json:"last_name" valid:"type(string),required~Last name cannot be empty"`
	Email        string    `json:"email" valid:"email, required~Email is required"`
	Password     string    `json:"password,omitempty" valid:"type(string),required~Password is required"`
	Gender       string    `json:"gender" valid:"type(string), required~Gender cannot be empty"`
	Address      string    `json:"address" valid:"type(string), required~Address cannot be empty"`
	MobileNumber string    `json:"mobile_number,omitempty" valid:"type(string), required~Mobile number cannot be empty"`
	Role         string    `gorm:"default: user" json:"role,omitempty" valid:"-"`
	Image        string    `json:"image" valid:"-"`
	CreatedAt    time.Time `json:"created_at"  valid:"-"`
	UpdatedAt    time.Time `json:"updated_at" valid:"-"`
}

type Users []User