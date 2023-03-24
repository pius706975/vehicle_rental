package command

import (
	"github.com/pius706975/backend/database/orm"
	libs "github.com/pius706975/backend/libs/server"
	"github.com/spf13/cobra"
)

var InitCommand = cobra.Command{
	Short: "go backend",
	Long:  `vehicle rental backend`,
}

func init() {
	InitCommand.AddCommand(libs.ServeCMD)
	InitCommand.AddCommand(orm.MigrateCMD)
	InitCommand.AddCommand(orm.SeedCMD)
}

func Run(args []string) error {
	InitCommand.SetArgs(args)

	return InitCommand.Execute()
}
