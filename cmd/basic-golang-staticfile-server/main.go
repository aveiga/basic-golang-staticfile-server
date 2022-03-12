package main

import (
	"github.com/aveiga/basic-golang-staticfile-server/internal/controllers"
	"github.com/aveiga/basic-golang-staticfile-server/internal/repositories"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/customdb"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db, err := customdb.GetDB()

	if err != nil {
		return
	}

	guitarRepo := repositories.NewGuitarRepo(db)
	h := controllers.NewBaseHandler(guitarRepo)

	router := gin.Default()

	router.POST("/guitars", h.CreateGuitar)
	router.GET("/guitars", h.GetGuitars)
	router.GET("/guitars/:id", controllers.SearchGuitars)
	router.DELETE("/guitars/:id", controllers.DeleteGuitar)

	router.Run(":8080")
}
