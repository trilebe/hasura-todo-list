package notification

import (
	"go-todo-app/database"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) save(notification *Notification) (*Notification, error) {
	err := r.db.Save(&notification).Error
	return notification, err
}

func NewRepository() *repository {
	db := database.Init()
	repository := repository{db}

	return &repository
}
