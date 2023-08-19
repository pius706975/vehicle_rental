package seeder

import "github.com/pius706975/backend/database/orm/models"

var UserSeed = models.Users{
	{
		// both password are the same, pius1234
		Name:         "Admin gahol",
		Username:     "admin123",
		Email:        "admin@gmail.com",
		Password:     "$2a$10$rYOxpbI1mgJyqPoiLw4BV.QuyGGN87u.5Yatm5mX6cuFtjdGwojJi",
		Gender:       "Male",
		Address:      "Ghost Cave Street",
		DateOfBirth:  "2023-03-12",
		MobileNumber: "085198762345",
		Role:         "admin",
		IsActive:     true,
	},

	{
		Name:         "Visitor gahol",
		Username:     "visitor123",
		Email:        "visitor@gmail.com",
		Password:     "$2a$10$EhHSap6LwdEwfxVHXBHjXedTekc4XngD6VfCy5VToOArRV9ZQpIFe",
		Gender:       "Male",
		Address:      "Ghost Cave Street",
		DateOfBirth:  "2023-03-12",
		MobileNumber: "085198762341",
		Role:         "user",
		IsActive:     true,
	},
}
