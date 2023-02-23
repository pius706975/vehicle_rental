package users

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"github.com/pius706975/backend/middleware"
)

func User(route *mux.Router, db *gorm.DB) {
	router := route.PathPrefix("/user").Subrouter()

	repo := NewUserRepo(db)
	service := NewUserService(repo)
	ctrl := NewUserCTRL(service)

	router.HandleFunc("", middleware.Handler(ctrl.GetAllUsers, middleware.AuthMiddle("admin"))).Methods("GET")

	router.HandleFunc("/register", ctrl.Register).Methods("POST")
	
	router.HandleFunc("/userprofile/edit", middleware.Handler(ctrl.UpdateUser, middleware.AuthMiddle("user", "admin"))).Methods("PUT")

	router.HandleFunc("/removeuser/{id}", middleware.Handler(ctrl.RemoveUser, middleware.AuthMiddle("admin"))).Methods("DELETE")
}
