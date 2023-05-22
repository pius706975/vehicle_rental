package history

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(route *mux.Router, db *gorm.DB) {
	router := route.PathPrefix("/history").Subrouter()

	repo := NewHistoryRepo(db)
	ctrl := NewHistoryCTRL(repo)

	router.HandleFunc("/", ctrl.GetAllHistory).Methods("GET")
	router.HandleFunc("/sort", ctrl.GetSortedHistory).Methods("GET")
}