package users

import (
	"github.com/gorilla/mux"
	"github.com/pius706975/backend/middleware"
	"gorm.io/gorm"
)

func New(route *mux.Router, db *gorm.DB) {
	router := route.PathPrefix("/user").Subrouter()

	repo := NewUserRepo(db)
	service := NewUserService(repo)
	ctrl := NewUserCTRL(service)

	router.HandleFunc("", middleware.Handler(ctrl.GetAllUsers, middleware.AuthMiddle("admin"))).Methods("GET")
	router.HandleFunc("/profile", middleware.Handler(ctrl.GetByID, middleware.AuthMiddle("admin", "user"))).Methods("GET")

	router.HandleFunc("/register", ctrl.Register).Methods("POST")

	router.HandleFunc("/profile/edit", middleware.Handler(ctrl.UpdateUser, middleware.AuthCloudUploadFile(), middleware.AuthMiddle("user", "admin"))).Methods("PUT")

	router.HandleFunc("/remove", middleware.Handler(ctrl.RemoveUser, middleware.AuthMiddle("admin", "user"))).Methods("DELETE")
}
