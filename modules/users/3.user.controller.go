package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/pius706975/backend/database/orm/models"
	"github.com/pius706975/backend/helper"
	"github.com/pius706975/backend/middleware"
)

type user_crtl struct {
	svc *user_service
}

func NewUserCTRL(svc *user_service) *user_crtl {
	return &user_crtl{svc}
}

// Register
func (c *user_crtl) Register(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.New(err.Error(), 500, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(&user)
	if err != nil {
		helper.New(err.Error(), 500, true).Send(w)
		return
	}

	c.svc.Register(&user).Send(w)
}

// Update user
func (c *user_crtl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-type", "application/json")

	

	userID, ok := r.Context().Value(middleware.UserID("user")).(uint)
	if !ok {
		helper.New("Unauthorized", 401, true).Send(w)
		return
	}

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.New(err.Error(), 400, true).Send(w)
		return
	}

	c.svc.UpdateUser(&user, uint(userID)).Send(w)
}

// REMOVE USER
func (c *user_crtl) RemoveUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		helper.New("Get the id first", 400, true).Send(w)
	}

	c.svc.RemoveUser(uint(id)).Send(w)
}

// GET ALL USERS
func (c *user_crtl) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	c.svc.GetAllUsers().Send(w)
}