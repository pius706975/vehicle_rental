package seeder

import "github.com/pius706975/backend/database/orm/models"

var CategorySeed = models.Categories{
	{
		CategoryID:   "8700e8d6-d113-4655-9e7e-e2741a0345d1",
		CategoryName: "Cars",
	},

	{
		CategoryID:   "7c2c2d26-f935-4f60-92b1-12aaad3a52d3",
		CategoryName: "Motorbikes",
	},

	{
		CategoryID:   "cd205e8a-6e8c-4d01-91e2-95d5c3a66fef",
		CategoryName: "Bicycles",
	},
}
