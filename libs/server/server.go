package libs

import (
	"log"
	"net/http"
	"os"

	"github.com/pius706975/backend/router"
	"github.com/spf13/cobra"
)

var ServeCMD = &cobra.Command{
	Use:   "serve",
	Short: "for running api server",
	RunE:  serve,
}

func serve(cmd *cobra.Command, args []string) error {
	mainRoute, err := router.RouterApp()
	if err != nil {
		return err
	}

	var address string = "0.0.0.0:3001"
	port := os.Getenv("PORT")
	if port != "" {
		address = ":" + port
	}

	log.Println("App is running on PORT", address)
	return http.ListenAndServe(address, mainRoute)
}