package category

import (
	"github.com/gorilla/mux"
	"github.com/pius706975/backend/middleware"
	"gorm.io/gorm"
)

func Category(route *mux.Router, db *gorm.DB) {
	router := route.PathPrefix("/vehicles/category").Subrouter()

	repo := NewCategoryRepo(db)
	service := NewCategoryService(repo)
	ctrl := NewCategoryCTRL(service)

	router.HandleFunc("", middleware.Handler(ctrl.GetAllCategories, middleware.AuthMiddle("admin"))).Methods("GET")
	router.HandleFunc("/addcategory", middleware.Handler(ctrl.AddCategory, middleware.AuthMiddle("admin"))).Methods("POST")
	router.HandleFunc("/removecategory/{id}", middleware.Handler(ctrl.RemoveCategory, middleware.AuthMiddle("admin"))).Methods("DELETE")
}