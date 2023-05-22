package users

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/schema"
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
		helper.New(err.Error(), 400, true).Send(w)
		return
	}

	c.svc.Register(&user).Send(w)
}

// Update user
func (c *user_crtl) UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	userID := r.Context().Value(middleware.UserID("user")).(string)

	var user models.User

	imageName := r.Context().Value("imageName").(string)
	user.Image = imageName

	err := schema.NewDecoder().Decode(&user, r.MultipartForm.Value)
	if err != nil {
		helper.New(err.Error(), 400, true).Send(w)
		return
	}

	c.svc.UpdateUser(&user, userID).Send(w)
}

// REMOVE USER
func (c *user_crtl) RemoveUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	user_id := r.Context().Value(middleware.UserID("user")).(string)

	c.svc.RemoveUser(user_id).Send(w)
}

// GET ALL USERS
func (c *user_crtl) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	c.svc.GetAllUsers().Send(w)
}

func (c *user_crtl) GetByID(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-type", "application/json")

	user_id := r.Context().Value(middleware.UserID("user")).(string)

	c.svc.GetByID(user_id).Send(w)
}