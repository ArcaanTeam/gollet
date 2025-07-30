package controllers

import (
	"fmt"
	"gollet/models"
	"gollet/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role"`
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Hash Password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	user := models.User{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: hashedPassword,
	}

	if err := uc.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not create user entry on database"})
		return
	}

	response := gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	}

	c.JSON(http.StatusCreated, response)
}

func (uc *UserController) GetProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	var user models.User
	result := uc.DB.First(&user, userID)
	if err := result.Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

func (uc *UserController) GetUsers(c *gin.Context) {
	var users []models.User
	if err := uc.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (uc *UserController) PromoteUser(c *gin.Context) {
	userToUpdateId := c.Param("id")
	if len(userToUpdateId) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "`id` param is not provided in url"})
		return
	}

	var payload struct {
		NewRole string `json:"new_role"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "'new_role' field is missing in body"},
		)
		return
	}

	var userToUpdate models.User
	if err := uc.DB.First(&userToUpdate, "id = ?", userToUpdateId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   fmt.Sprintf("Failed to find user with id: %s", userToUpdateId),
			"details": err.Error(),
		})
	}

	if err := uc.DB.Model(&userToUpdate).Update("role", payload.NewRole).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   fmt.Sprintf("Failed to update user %s", userToUpdateId),
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":    userToUpdate.ID,
			"name":  userToUpdate.Name,
			"email": userToUpdate.Email,
			"role":  userToUpdate.Role,
		},
	})
}
