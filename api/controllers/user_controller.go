package controllers

import (
	"gollet/api/payloads"
	"gollet/services"
	"gollet/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
}

type userController struct {
	s services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userController{s: service}
}

func (c *userController) CreateUser(ctx *gin.Context) {
	var input payloads.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "malformed request body",
			"details": err.Error(),
		})
	}

	createdUser, err := c.s.CreateUser(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "unable to create user",
			"details": err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	userIdString, _, err := utils.GetIdParam(ctx)
	if err != nil {
		return
	}

	var input payloads.UpdateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	updatedUser, err := c.s.UpdateUser(userIdString, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "unable to update user",
			"details": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, updatedUser)
}

func (c *userController) GetUserByID(ctx *gin.Context) {
	idString, _, err := utils.GetIdParam(ctx)
	if err != nil {
		return
	}

	user, err := c.s.GetUserByID(idString)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   "user not found",
				"details": "",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "internal server error",
				"details": err.Error(),
			})
		}
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *userController) GetProfile(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	userIdString, ok := userID.(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid user id",
			"details": "",
		})
	}
	user, err := c.s.GetUserByID(userIdString)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   "user not found",
				"details": "",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "internal server error",
				"details": err.Error(),
			})
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

func (c *userController) GetUsers(ctx *gin.Context) {
	users, err := c.s.GetUsers()
	if err != nil {
		// TODO: handle not-found branch separately
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
