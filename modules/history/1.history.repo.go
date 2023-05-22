package history

import (
	"github.com/pius706975/backend/database/orm/models"
	"gorm.io/gorm"
)

type history_repo struct {
	db *gorm.DB
}

func NewHistoryRepo(db *gorm.DB) *history_repo {
	return &history_repo{db}
}

func (r *history_repo) GetAllHistory() (*models.Histories, error) {
	var data models.Histories

	result := r.db.Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func (r *history_repo) GetSortedHistory() (*models.Histories, error) {
	var data models.Histories

	result := r.db.Order("created_at desc").Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}