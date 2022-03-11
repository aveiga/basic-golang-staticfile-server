package guitars

import (
	"fmt"
	"github.com/aveiga/basic-golang-staticfile-server/domain/guitars"
	guitars2 "github.com/aveiga/basic-golang-staticfile-server/services/guitars"
	"github.com/aveiga/basic-golang-staticfile-server/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateGuitar(c *gin.Context) {
	var guitar guitars.Guitar
	if err := c.ShouldBindJSON(&guitar); err != nil {
		error := errors.RestError {
			Message: "Invalid format",
			Status: http.StatusBadRequest,
			Code: "bad_request",
		}
		c.JSON(error.Status, error)
		return
	}
	fmt.Println(guitar)
	result, saveError := guitars2.CreateGuitar(guitar)
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
