package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aveiga/basic-golang-staticfile-server/internal/services"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/models"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/customerrors"
	"github.com/gin-gonic/gin"
)

type BaseHandler struct {
	guitarRepo models.GuitarRepository
}

func NewBaseHandler(guitarRepo models.GuitarRepository) *BaseHandler {
	return &BaseHandler{
		guitarRepo: guitarRepo,
	}
}

func (h *BaseHandler) CreateGuitar(c *gin.Context) {
	var guitar models.Guitar
	if err := c.ShouldBindJSON(&guitar); err != nil {
		error := customerrors.RestError{
			Message: "Invalid format",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
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

func (h *BaseHandler) GetGuitars(c *gin.Context) {
	guitars, err := h.guitarRepo.FindAll()
	if err != nil {
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, guitars)
}

func SearchGuitars(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func DeleteGuitar(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
