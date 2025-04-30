package service

import (
	"errors"
	"fmt"
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

func DeactivateUser(id uint) error {
	var user model.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return errors.New("user not found")
	}
	user.Active = false
	return config.DB.Save(&user).Error
}

func ReactivateUser(id uint) error {
	var user model.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return errors.New("user not found")
	}

	// Jika user sudah aktif, tidak perlu dilakukan apa-apa
	if user.Active {
		return errors.New("user is already active")
	}

	// Aktifkan kembali user
	user.Active = true
	return config.DB.Save(&user).Error
}
func SearchUsers(filters map[string]string) ([]model.User, error) {
	var users []model.User
	query := config.DB.Model(&model.User{})

	// Print isi filters
	fmt.Println("Filters received:", filters)

	for key, value := range filters {
		if value == "" {
			continue
		}

		switch key {
		case "full_name", "email":
			query = query.Where(fmt.Sprintf("%s ILIKE ?", key), "%"+value+"%")
		case "role", "status":
			query = query.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}

	err := query.Find(&users).Error
	return users, err
}
func UpdateUserPassword(userID uint, newPassword string) error {
	if len(newPassword) < 6 {
		return errors.New("password harus minimal 8 karakter")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	result := config.DB.Model(&model.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"password": string(hashedPassword),
		})

	if result.RowsAffected == 0 {
		return errors.New("user tidak ditemukan")
	}

	return result.Error
}

func CountActiveAdmins() (int64, error) {
	var count int64 // Change from int to int64
	err := config.DB.Model(&model.User{}).Where("role = ? AND active = ?", "admin", true).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
