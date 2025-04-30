package controller

import (
	"fmt"
	"go-gin-auth/model"
	"go-gin-auth/service"
	"go-gin-auth/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("PPL-K4-2025")

func Register(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid input", err.Error(), nil)
		return
	}

	if err := service.CreateUser(&user); err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to create user", err.Error(), nil)
		return
	}

	// Menambahkan aktivitas log setelah user berhasil didaftarkan
	err := service.LogActivity(user.ID, user.FullName, "Register", "User registered successfully.", c)
	if err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to log activity", err.Error(), nil)
		return
	}
	utils.Respond(c, http.StatusCreated, "User registered successfully", nil, user)
}

func Login(c *gin.Context) {
	var input model.User

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid request", err.Error(), nil)
		return
	}

	user, err := service.GetUserByEmail(input.Email)
	if err != nil {
		utils.Respond(c, http.StatusUnauthorized, "Login failed", "User not found", nil)
		return
	}

	// Mengecek jika user tidak aktif
	if !user.Active {
		utils.Respond(c, http.StatusUnauthorized, "Login failed", "User is inactive", nil)
		return
	}

	match := service.VerifyPassword(input.Password, user.Password)

	if !match {
		utils.Respond(c, http.StatusUnauthorized, "Login failed", "Incorrect password", nil)
		return
	}

	// generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"role":      user.Role, // ← penting untuk middleware
		"full_name": user.FullName,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Could not generate token", err.Error(), nil)
		return
	}

	// Catat log aktivitas login
	err = service.LogActivity(user.ID, user.FullName, "Login", "User logged in successfully", c)
	if err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to log activity", err.Error(), nil)
		return
	}

	utils.Respond(c, http.StatusOK, "Login successful", nil, gin.H{
		"token": tokenString,
	})
}

func Logout(c *gin.Context) {
	// Biasanya di sisi frontend: hapus token dari storage.
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}

func GetUsers(c *gin.Context) {
	// // Cek nilai di context
	// userID, userIDExists := c.Get("user_id")
	// if !userIDExists {
	// 	utils.Respond(c, http.StatusUnauthorized, "Unauthorized", "Missing user_id in context", nil)
	// 	return
	// }

	// // Cek nilai user_id
	// fmt.Printf("User ID: %v\n", userID)
	users, err := service.GetAllUsers()
	if err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to get users", err.Error(), nil)
		return
	}
	utils.Respond(c, http.StatusOK, "Users retrieved successfully", nil, users)
}

func GetUser(c *gin.Context) {
	var id uint
	fmt.Sscanf(c.Param("id"), "%d", &id)
	user, err := service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUserOri(c *gin.Context) {
	var id uint
	fmt.Sscanf(c.Param("id"), "%d", &id)
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.UpdateUser(id, user)
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

func DeleteUserOri(c *gin.Context) {
	var id uint
	fmt.Sscanf(c.Param("id"), "%d", &id)
	service.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func DeleteUser(c *gin.Context) {
	var id uint
	if _, err := fmt.Sscanf(c.Param("id"), "%d", &id); err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid ID parameter", err.Error(), nil)
		return
	}
	// Dapatkan informasi user yang akan dihapus
	targetUser, err := service.GetUserByID(id)
	if err != nil {
		utils.Respond(c, http.StatusNotFound, "User not found", err.Error(), nil)
		return
	}

	err = service.DeleteUser(targetUser.ID)
	if err != nil {
		utils.Respond(c, http.StatusNotFound, "User not found", err.Error(), nil)
		return
	}
	// Ambil user yang sedang login dari context (misalnya dari middleware)
	currentUserIDFloat, _ := c.Get("user_id")
	currentUserID := uint(currentUserIDFloat.(float64))
	currentFullName, _ := c.Get("full_name")

	fmt.Printf("Logged-in user: %v (ID %v)\n", currentFullName, currentUserID)

	// Catat log
	description := fmt.Sprintf("User %s (ID %d) deleted user %s (ID %d)",
		currentFullName, currentUserID, targetUser.FullName, targetUser.ID)

	_ = service.LogActivity(currentUserID, currentFullName.(string), "DeleteUser", description, c)

	utils.Respond(c, http.StatusOK, "User deleted successfully", nil, nil)
}

func UpdateUser(c *gin.Context) {
	var id uint
	if _, err := fmt.Sscanf(c.Param("id"), "%d", &id); err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid ID parameter", err.Error(), nil)
		return
	}

	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid request body", err.Error(), nil)
		return
	}

	// Cek apakah user dengan ID tersebut ada
	existingUser, err := service.GetUserByID(id)
	if err != nil {
		utils.Respond(c, http.StatusNotFound, "User not found", err.Error(), nil)
		return
	}

	// Update field yang ingin diubah (hindari overwrite ID/Password langsung)
	existingUser.Email = input.Email
	existingUser.FullName = input.FullName
	existingUser.Role = input.Role

	if err := service.UpdateUser(id, existingUser); err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to update user", err.Error(), nil)
		return
	}

	utils.Respond(c, http.StatusOK, "User updated successfully", nil, existingUser)
}

func DeactivateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid user ID", nil, err.Error())
		return
	}

	// Ambil user terlebih dahulu
	user, err := service.GetUserByID(uint(id))
	if err != nil {
		utils.Respond(c, http.StatusNotFound, "User not found", nil, err.Error())
		return
	}

	// Jika sudah nonaktif
	if !user.Active {
		utils.Respond(c, http.StatusBadRequest, "User is already deactivated", nil, nil)
		return
	}

	// Cek jumlah admin aktif lainnya
	activeAdminsCount, err := service.CountActiveAdmins()
	if err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Error", "Failed to count active admins", nil)
		return
	}

	// Jika hanya satu admin yang aktif, tolak permintaan
	if activeAdminsCount <= 1 {
		utils.Respond(c, http.StatusForbidden, "Forbidden", "You cannot deactivate your account, at least one admin must be active", nil)
		return
	}

	err = service.DeactivateUser(uint(id))
	if err != nil {
		utils.Respond(c, http.StatusNotFound, "User not found", nil, err.Error())
		return
	}

	// // Tambahkan log aktivitas
	// service.CreateUserLog(uint(id), "Deactivate User")

	utils.Respond(c, http.StatusOK, "User deactivated successfully", nil, nil)
}
func ReactivateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid user ID", nil, err.Error())
		return
	}

	err = service.ReactivateUser(uint(id))
	if err != nil {
		utils.Respond(c, http.StatusBadRequest, "Failed to reactivate user", nil, err.Error())
		return
	}

	utils.Respond(c, http.StatusOK, "User reactivated successfully", nil, nil)
}

func SearchUsers(c *gin.Context) {
	filters := map[string]string{
		"full_name": c.Query("full_name"),
		"email":     c.Query("email"),
		"role":      c.Query("role"),
	}

	users, err := service.SearchUsers(filters)
	if err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Search failed", nil, err.Error())
		return
	}

	utils.Respond(c, http.StatusOK, "Users fetched", nil, users)
}
func ResetUserPassword(c *gin.Context) {
	var id uint
	if _, err := fmt.Sscanf(c.Param("id"), "%d", &id); err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid ID parameter", err.Error(), nil)
		return
	}

	var body struct {
		NewPassword string `json:"new_password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid request body", err.Error(), nil)
		return
	}

	if body.NewPassword == "" {
		utils.Respond(c, http.StatusBadRequest, "Password is required", "new_password is empty", nil)
		return
	}

	err := service.UpdateUserPassword(id, body.NewPassword)
	if err != nil {
		utils.Respond(c, http.StatusBadRequest, "Failed to update password", err.Error(), nil)
		return
	}

	utils.Respond(c, http.StatusOK, "Password updated successfully", nil, nil)
}
