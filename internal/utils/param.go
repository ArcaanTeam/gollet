package utils

import (
	"errors"
	"gollet/internal/constants"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIdParam(ctx *gin.Context) (string, int, error) {
	idString := ctx.Param("id")
	if idString == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id param is required"})
		return "", -1, errors.New(constants.ErrIDParamNotProvided)
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return idString, -1, errors.New(constants.ErrInvalidIDParam)
	}
	return idString, id, nil
}
