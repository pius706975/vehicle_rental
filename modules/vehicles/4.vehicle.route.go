package vehicles

import (
	"github.com/gorilla/mux"
	"github.com/pius706975/backend/middleware"
	"gorm.io/gorm"
)

func New(route *mux.Router, db *gorm.DB) {
	router := route.PathPrefix("/vehicles").Subrouter()

	repo := NewVehicleRepo(db)
	service := NewVehicleService(repo)
	ctrl := NewVehicleCTRL(service)

	router.HandleFunc("/", ctrl.GetAllVehicles).Methods("GET")
	router.HandleFunc("/popular/sort", ctrl.GetPopularVehicle).Methods("GET")
	router.HandleFunc("/category/{category}", ctrl.GetVehicleByCategory).Methods("GET")
	router.HandleFunc("/model/{model}", ctrl.GetVehicleByModel).Methods("GET")

	router.HandleFunc("/addvehicle", middleware.Handler(ctrl.AddNewVehicle, middleware.AuthMiddle("admin"))).Methods("POST")
	router.HandleFunc("/removevehicle/{id}", middleware.Handler(ctrl.RemoveVehicle, middleware.AuthMiddle("admin"))).Methods("DELETE")
	router.HandleFunc("/updatevehicle/{id}", middleware.Handler(ctrl.UpdateVehicle, middleware.AuthMiddle("admin"))).Methods("PUT")
}