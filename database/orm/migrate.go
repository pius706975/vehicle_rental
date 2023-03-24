package orm

import (
	"log"

	"github.com/pius706975/backend/database/orm/models"
	"github.com/spf13/cobra"
)

var MigrateCMD = &cobra.Command{
	Use:   "migrate",
	Short: "db migration",
	RunE:  dbMigrate,
}

var migUp bool
var migDown bool

func init() {
	MigrateCMD.Flags().BoolVarP(&migUp, "dbUP", "u", true, "run migration up")
	
	MigrateCMD.Flags().BoolVarP(&migDown, "dbDOWN", "d", false, "run migration down")
}

func dbMigrate(cmd *cobra.Command, args []string) error {

	db, err := NewDB()
	if err != nil {
		return err
	}

	if migDown {
		log.Println("Migration down done")
		return db.Migrator().DropTable(&models.User{}, &models.Category{}, &models.Vehicle{}, &models.Reservation{}, &models.History{})
	}

	if migUp {
		log.Println("Migration up done")
		return db.AutoMigrate(&models.User{}, &models.Category{}, &models.Vehicle{}, &models.Reservation{}, &models.History{})
	}

	return nil
}