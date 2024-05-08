package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CreateUser создает нового пользователя в базе данных.
func CreateUser(db *gorm.DB, username, password string) (*User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username: username,
		Password: hashedPassword,
	}

	// Сохраняем пользователя в базе данных
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// hashPassword хеширует пароль пользователя.
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
