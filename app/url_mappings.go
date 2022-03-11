package app

import (
	"github.com/aveiga/basic-golang-staticfile-server/controllers/guitars"
)

func mapUrls() {
	router.POST("/guitars", guitars.CreateGuitar)
	router.GET("/guitars", guitars.GetGuitars)
	router.GET("/guitars/:id", guitars.SearchGuitars)
	router.DELETE("/guitars/:id", guitars.DeleteGuitar)
}
