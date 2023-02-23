package category

import (
	"encoding/json"
	"net/http"
	"strconv"

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

	id_ := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(id_, 10, 64)
	if err != nil {
		helper.New("Get the ID first", 400, true).Send(w)
		return
	}

	c.svc.RemoveCategory(uint(id)).Send(w)
}
// func (c *category_ctrl) RemoveCategory(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json")

// 	var category models.Category

// 	// parse id from request
// 	idStr := mux.Vars(r)["id"]
// 	id, err := strconv.ParseUint(idStr, 10, 64)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	err = c.repo.db.Where("category_id = ?", id).First(&category).Error
// 	if err != nil {
// 		if category.CategoryID == 0 {
// 			respond := helper.New("Category not found", 400, true)
// 			respond.Send(w)
// 			return
// 		}
// 	}

// 	// err = c.repo.db.Unscoped().Where("category_id = ?", id).Delete(&category).Error
// 	// if err != nil {
// 	// 	http.Error(w, err.Error(), http.StatusBadRequest)
// 	// 	return
// 	// }

// 	err = c.repo.RemoveCategory(uint(id))

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(map[string]string{"Message": "Category has been deleted"})
// }