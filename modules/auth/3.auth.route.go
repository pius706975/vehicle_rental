package auth

import (
	"github.com/gorilla/mux"
	"github.com/pius706975/backend/modules/users"
	"gorm.io/gorm"
)

func New(route *mux.Router, db *gorm.DB) {
	router := route.PathPrefix("/auth").Subrouter()

	repo := users.NewUserRepo(db)
	service := NewAuthService(repo)
	ctrl := NewAuthCtrl(*service)

	router.HandleFunc("/login", ctrl.Login).Methods("POST")

	router.HandleFunc("/verify_email/{token}", ctrl.VerifyEmail).Methods("GET")
}
