package main

import (
	"log"
	// "net/http"
	"os"

	"github.com/asaskevich/govalidator"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pius706975/backend/command"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {

	err := command.Run(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// r, err := router.RouterApp()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("App is running on PORT 3001")
	// err = http.ListenAndServe(":3001", r)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
