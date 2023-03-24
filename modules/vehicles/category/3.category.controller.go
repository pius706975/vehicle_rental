package category

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/pius706975/backend/database/orm/models"
	"github.com/pius706975/backend/helper"
)

type category_ctrl struct {
	svc *category_service
}

func NewCategoryCTRL(svc *category_service) *category_ctrl {
	return &category_ctrl{svc}
}

// ADD CATEGORY
func (c *category_ctrl) AddCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	var category models.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		helper.New(err.Error(), 500, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(&category)
	if err != nil {
		helper.New(err.Error(), 500, true).Send(w)
		return
	}

	c.svc.AddCategory(&category).Send(w)
}

// GET ALL CATEGORIES
func (c *category_ctrl) GetAllCategories(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	c.svc.GetAllCategories().Send(w)
}

// REMOVE CATEGORY
func (c *category_ctrl) RemoveCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	idStr := mux.Vars(r)
	id, ok := idStr["id"]
	if !ok {
		helper.New("Get the ID first", 400, true).Send(w)
		return
	}

	c.svc.RemoveCategory(id).Send(w)
}