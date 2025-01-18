package user

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             string `gorm:"type:uuid;primaryKey" json:"id"`
	UserName       string `gorm:"type:text" json:"username"`
	HashedPassword string `gorm:"type:text"`
}

func NewUser(username string, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	return &User{UserName: username, HashedPassword: string(hashedPassword)}, nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(providedPassword))
	return err
}
