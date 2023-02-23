package orm

import (
	"github.com/pius706975/backend/database/orm/models"
	"github.com/spf13/cobra"
)

var MigrateCMD = &cobra.Command{
	Use:   "migrate",
	Short: "for running database migration",
	RunE:  dbMigrate,
}

var migUP bool
var migDOWN bool

func init() {
	MigrateCMD.Flags().BoolVarP(&migUP, "dbUP", "u", true, "Running auto migration")

	MigrateCMD.Flags().BoolVarP(&migDOWN, "dbDOWN", "d", false, "Running auto reset migration")
}

func dbMigrate(cmd *cobra.Command, args []string) error {
	db, err := NewDB()
	if err != nil {
		return err
	}

	if migUP {
		return db.AutoMigrate(&models.User{}, &models.Vehicle{}, &models.Category{}, &models.History{}, &models.Reservation{})
	}

	if migDOWN {
		return db.Migrator().DropTable(&models.User{}, &models.Vehicle{}, &models.Category{}, &models.History{}, &models.Reservation{})
	}

	return nil
}
