package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pius706975/backend/database/orm"
	"github.com/pius706975/backend/modules/auth"
	"github.com/pius706975/backend/modules/history"
	"github.com/pius706975/backend/modules/reservation"
	"github.com/pius706975/backend/modules/users"
	"github.com/pius706975/backend/modules/vehicles"
	"github.com/pius706975/backend/modules/vehicles/category"
)

func RouterApp() (*mux.Router, error) {

	mainRoute := mux.NewRouter()

	db, err := orm.NewDB()

	if err != nil {
		return nil, err
	}

	subRouter := mainRoute.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/", homeHandler).Methods("GET")

	// mainRoute.HandleFunc("/", homeHandler).Methods("GET")

	vehicles.New(subRouter, db)
	users.New(subRouter, db)
	auth.New(subRouter, db)
	reservation.New(subRouter, db)
	history.New(subRouter, db)
	category.New(subRouter, db)

	// mainRoute.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "text/html")
	// 	w.Write([]byte("Hello API"))
	// })

	return mainRoute, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is Vehicle Rental backend!"))
}
