package router

import (
	"github.com/gorilla/mux"
	"github.com/pius706975/backend/database/orm"
	"github.com/pius706975/backend/modules/history"
	"github.com/pius706975/backend/modules/reservation"
	"github.com/pius706975/backend/modules/users"
	"github.com/pius706975/backend/modules/auth"
	"github.com/pius706975/backend/modules/vehicles"
	"github.com/pius706975/backend/modules/vehicles/category"
)

func RouterApp() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	db, err := orm.NewDB()
	
	if err != nil {
		return nil, err
	}

	vehicles.New(mainRoute, db)
	users.User(mainRoute, db)
	auth.Auth(mainRoute, db)
	reservation.Reserve(mainRoute, db)
	history.History(mainRoute, db)
	category.Category(mainRoute, db )


	return mainRoute, nil
}