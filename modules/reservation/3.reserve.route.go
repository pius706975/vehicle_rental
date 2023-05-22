package reservation

import (
	"github.com/gorilla/mux"
	"github.com/pius706975/backend/middleware"
	"gorm.io/gorm"
)

func New(route *mux.Router, db *gorm.DB) {
	router := route.PathPrefix("/reservation").Subrouter()

	repo := NewReserveRepo(db)
	// service := NewReserveService(repo)
	ctrl := NewReserveCTRL(repo)

	router.HandleFunc("/reserve", middleware.Handler(ctrl.CreateReservation, middleware.AuthMiddle("user"))).Methods("POST")

	router.HandleFunc("/return", middleware.Handler(ctrl.ReturnVehicle, middleware.AuthMiddle("user"))).Methods("PUT")

	router.HandleFunc("/payment", middleware.Handler(ctrl.Payment, middleware.AuthMiddle("user"))).Methods("POST")
}