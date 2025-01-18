package user

import (
	"go-todo-app/database"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) save(user *User) (*User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *repository) findByUsername(username string) (*User, error) {
	var user User
	if err := r.db.Where(&User{UserName: username}).First(&user).Error; err != nil {
		return nil, &Errors.NotFound
	}
	return &user, nil
}

func NewRepository() *repository {
	db := database.Init()
	repository := repository{db}
	db.AutoMigrate(&User{})

	return &repository
}
