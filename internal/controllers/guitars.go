package controllers

import (
	"fmt"
	"net/http"

	"github.com/aveiga/basic-golang-staticfile-server/internal/services"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/models"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/customerrors"
	"github.com/gin-gonic/gin"
)

func CreateGuitar(c *gin.Context) {
	var guitar models.Guitar
	if err := c.ShouldBindJSON(&guitar); err != nil {
		error := customerrors.RestError {
			Message: "Invalid format",
			Status: http.StatusBadRequest,
			Code: "bad_request",
		}
		c.JSON(error.Status, error)
		return
	}
	fmt.Println(guitar)
	result, saveError := services.CreateGuitar(guitar)
	if saveError != nil {
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetGuitars(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func SearchGuitars(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func DeleteGuitar(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
