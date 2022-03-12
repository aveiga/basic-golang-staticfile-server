package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/aveiga/basic-golang-staticfile-server/internal/controllers"
	"github.com/aveiga/basic-golang-staticfile-server/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun/driver/pgdriver"
)

func main() {
	godotenv.Load()
	// ctx := context.Background()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_ADDRESS"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	// dsn := "unix://user:pass@dbname/var/run/postgresql/.s.PGSQL.5432"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// db := bun.NewDB(sqldb, pgdialect.New())

	guitarRepo := repositories.NewGuitarRepo(sqldb)
	h := controllers.NewBaseHandler(guitarRepo)

	// guitars := make([]models.Guitar, 0)
	// err := db.NewSelect().
	// 	Model(&guitars).
	// 	Scan(ctx)
	// fmt.Printf("all users: %v\n\n", guitars)

	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	router := gin.Default()

	router.POST("/guitars", h.CreateGuitar)
	router.GET("/guitars", h.GetGuitars)
	router.GET("/guitars/:id", controllers.SearchGuitars)
	router.DELETE("/guitars/:id", controllers.DeleteGuitar)

	router.Run(":8080")
}
