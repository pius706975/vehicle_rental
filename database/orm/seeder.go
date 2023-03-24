package orm

import (
	"fmt"
	"log"
	
	seeder "github.com/pius706975/backend/database/orm/seeders"

	"github.com/pius706975/backend/database/orm/models"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type seederData struct {
	name  string
	model interface{}
	size  int
}

var SeedCMD = &cobra.Command{
	Use: "seed",
	Short: "For running db seeder",
	RunE: Seed,
}

var seedUP bool
var seedDOWN bool

func init()  {
	SeedCMD.Flags().BoolVarP(&seedUP, "seedUP", "u", true, "run seed up")

	SeedCMD.Flags().BoolVarP(&seedDOWN, "seedDOWN", "d", false, "run seed down")
}

func seedDown(db *gorm.DB) error {
	
	var err error

	var seedModel = []seederData{
		{
			name: models.User{}.TableName(),
			model: models.User{},
		},

		{
			name: models.Category{}.TableName(),
			model: models.Category{},
		},

		{
			name: models.Vehicle{}.TableName(),
			model: models.Vehicle{},
		},
	}

	for _, data := range seedModel {
		log.Println("Delete seeding data for ", data.name)
		sql := fmt.Sprintf("DELETE FROM %v ", data.name)
		err = db.Exec(sql).Error
	}

	return err
}

func seedUp(db *gorm.DB) error {
	
	var err error

	var seedModel = []seederData{
		{
			name: "user",
			model: seeder.UserSeed,
			size: cap(seeder.UserSeed),
		},

		{
			name: "category",
			model: seeder.CategorySeed,
			size: cap(seeder.CategorySeed),
		},

		{
			name: "vehicle",
			model: seeder.VehicleSeed,
			size: cap(seeder.VehicleSeed),
		},
	}

	for _, data := range seedModel {

		log.Println("create seeding data for ", data.name)
		err = db.CreateInBatches(data.model, data.size).Error
	}

	return err
}

func Seed(cmd *cobra.Command, args []string) error {
	
	var err error

	db, err := NewDB()
	if err != nil {
		return err
	}

	if seedDOWN {
		err = seedDown(db)
		return err
	}

	if seedUP {
		err = seedUp(db)
	}

	return err
}