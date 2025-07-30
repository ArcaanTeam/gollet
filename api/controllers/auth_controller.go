package controllers

import (
	"gollet/api/payloads"
	"gollet/constants"
	"gollet/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
	s services.AuthService
}

func NewAuthController(service services.AuthService) *authController {
	return &authController{s: service}
}

func (c *authController) Login(ctx *gin.Context) {
	var input payloads.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := c.s.Login(input)
	if err != nil {
		var statusCode int
		errorString := err.Error()

		switch errorString {
		case constants.ErrAuthUserNotFound:
			statusCode = http.StatusNotFound
		case constants.ErrAuthUnauthorized:
			statusCode = http.StatusUnauthorized
		case constants.ErrAuthGenerateTokenFailed:
			statusCode = http.StatusInternalServerError
		}
		ctx.JSON(
			statusCode,
			gin.H{
				"error": errorString,
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
			"role":  user.Role,
		},
	})
}
