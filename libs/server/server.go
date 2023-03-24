package libs

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pius706975/backend/router"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var ServeCMD = &cobra.Command{
	Use:   "serve",
	Short: "for running api server",
	RunE:  serve,
}

func corsHandler() *cors.Cors {
	
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},

		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	return c
}

func serve(cmd *cobra.Command, args []string) error {

	mainRoute, err := router.RouterApp()
	if err != nil {
		return err
	}

	var address string = "0.0.0.0:3001"
	if PORT := os.Getenv("PORT"); PORT != "" {
		address = "127.0.0.1:" + PORT
	}

	cors := corsHandler()

	serve := &http.Server{
		Addr: address,
		WriteTimeout: time.Second,
		ReadTimeout: time.Second,
		IdleTimeout: time.Minute,
		Handler: cors.Handler(mainRoute),
	}

	log.Println("App is running on PORT 3001")

	return serve.ListenAndServe()
}