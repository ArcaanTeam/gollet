package controllers

import (
	"gollet/internal/constants"
	"gollet/internal/interfaces/payloads"
	"gollet/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authController struct {
	s usecases.AuthService
}

func NewAuthController(service usecases.AuthService) *authController {
	return &authController{s: service}
}

func (c *authController) Login(ctx *gin.Context) {
	var input payloads.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := c.s.Login(input)
	// FIXME: move error handling to error middleware
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
