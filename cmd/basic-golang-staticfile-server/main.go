package main

import (
	"github.com/aveiga/basic-golang-staticfile-server/internal/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/guitars", controllers.CreateGuitar)
	router.GET("/guitars", controllers.GetGuitars)
	router.GET("/guitars/:id", controllers.SearchGuitars)
	router.DELETE("/guitars/:id", controllers.DeleteGuitar)
	
	router.Run(":8080")
}
