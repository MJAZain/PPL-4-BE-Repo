package controller

import (
	"fmt"
	"go-gin-auth/model"
	"go-gin-auth/service"
	"go-gin-auth/utils"
	"net/http"
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

	match := service.VerifyPassword(input.Password, user.Password)

	if !match {
		utils.Respond(c, http.StatusUnauthorized, "Login failed", "Incorrect password", nil)
		return
	}

	// generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Could not generate token", err.Error(), nil)
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
	users, _ := service.GetAllUsers()
	c.JSON(http.StatusOK, users)
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

	err := service.DeleteUser(id)
	if err != nil {
		utils.Respond(c, http.StatusNotFound, "User not found", err.Error(), nil)
		return
	}

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
