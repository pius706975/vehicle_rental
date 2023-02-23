package auth

import (
	"github.com/gorilla/mux"
	"github.com/pius706975/backend/modules/users"
	"gorm.io/gorm"
)

func Auth(route *mux.Router, db *gorm.DB) {
	router := route.PathPrefix("/user/auth").Subrouter()

	repo := users.NewUserRepo(db)
	service := NewAuthService(repo)
	ctrl := NewAuthCtrl(*service)

	router.HandleFunc("", ctrl.Login).Methods("POST")
}
