package auth

import (
	"encoding/json"
	"net/http"

	"github.com/pius706975/backend/database/orm/models"
	"github.com/pius706975/backend/helper"
)

type auth_ctrl struct {
	repo auth_service
}

func NewAuthCtrl(repo auth_service) *auth_ctrl {
	return &auth_ctrl{repo}
}

func (c *auth_ctrl) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") 
	
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.New(err.Error(), 500, true).Send(w)
		return
	}

	c.repo.Login(&user).Send(w)
}