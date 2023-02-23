package history

import (
	"encoding/json"
	"net/http"
)

type history_ctrl struct {
	repo *history_repo
}

func NewHistoryCTRL(repo *history_repo) *history_ctrl {
	return &history_ctrl{repo}
}

func (c *history_ctrl) GetAllHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	result, err := c.repo.GetAllHistory()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(result)
}

func (c *history_ctrl) GetSortedHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	result, err := c.repo.GetSortedHistory()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(result)
}