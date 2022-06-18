package controllers

import (
	"net/http"

	"github.com/aveiga/basic-golang-staticfile-server/pkg/models"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/customerrors"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/customlogger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GuitarController struct {
	guitarService models.GuitarService
	logger        *zap.SugaredLogger
}

func NewGuitarController(guitarService models.GuitarService) *GuitarController {
	return &GuitarController{
		guitarService: guitarService,
		logger:        customlogger.NewCustomLogger(),
	}
}

func (gc *GuitarController) CreateGuitar(c *gin.Context) {
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

	result, saveError := gc.guitarService.CreateGuitar(&guitar)
	if saveError != nil {
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (gc *GuitarController) GetGuitars(c *gin.Context) {
	guitars, err := gc.guitarService.GetGuitars()
	if err != nil {
		gc.logger.Fatal(err.Error())
		return
	}

	c.JSON(http.StatusOK, guitars)
}

func (gc *GuitarController) SearchGuitars(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func (gc *GuitarController) DeleteGuitar(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
