package service

import (
	"errors"
	"go-gin-auth/config"
	"go-gin-auth/model"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash for the given password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(user *model.User) error {
	hashedPassword, _ := HashPassword(user.Password)
	user.Password = hashedPassword
	return config.DB.Create(user).Error
}

func GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := config.DB.Find(&users).Error
	return users, err
}

func GetUserByEmail(email string) (model.User, error) {
	var user model.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func GetUserByID(id uint) (model.User, error) {
	var user model.User
	err := config.DB.First(&user, id).Error
	return user, err
}

func UpdateUser(id uint, updated model.User) error {
	return config.DB.Model(&model.User{}).Where("id = ?", id).Updates(updated).Error
}

func DeleteUser(id uint) error {
	result := config.DB.Delete(&model.User{}, id)
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return result.Error
}
