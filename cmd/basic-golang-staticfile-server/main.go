package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aveiga/basic-golang-staticfile-server/internal/controllers"
	"github.com/aveiga/basic-golang-staticfile-server/internal/repositories"
	"github.com/aveiga/basic-golang-staticfile-server/internal/services"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/customdb"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// ch, _ := customamqp.GetAMQPChannel()
	// q, _ := ch.QueueDeclare(
	// 	"hello", // name
	// 	false,   // durable
	// 	false,   // delete when unused
	// 	false,   // exclusive
	// 	false,   // no-wait
	// 	nil,     // arguments
	// )

	// body := "Hello World!"
	// _ = ch.Publish(
	// 	"",     // exchange
	// 	q.Name, // routing key
	// 	false,  // mandatory
	// 	false,  // immediate
	// 	amqp.Publishing{
	// 		ContentType: "text/plain",
	// 		Body:        []byte(body),
	// 	})

	// msgs, _ := ch.Consume(
	// 	q.Name, // queue
	// 	"",     // consumer
	// 	true,   // auto-ack
	// 	false,  // exclusive
	// 	false,  // no-local
	// 	false,  // no-wait
	// 	nil,    // args
	// )

	// forever := make(chan bool)

	// go func() {
	// 	for d := range msgs {
	// 		log.Printf("Received a message: %s", d.Body)
	// 	}
	// }()

	// log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	// <-forever

	ctx := context.Background()
	db, dbErr := customdb.GetDB()

	if dbErr != nil {
		log.Fatal("Error creating database connection")
		return
	}

	guitarRepo := repositories.NewGuitarRepo(db, ctx)
	guitarService := services.NewGuitarService(guitarRepo)
	guitarController := controllers.NewGuitarController(guitarService)

	router := gin.Default()

	router.POST("/guitars", guitarController.CreateGuitar)
	router.GET("/guitars", guitarController.GetGuitars)
	router.GET("/guitars/:id", guitarController.SearchGuitars)
	router.DELETE("/guitars/:id", guitarController.DeleteGuitar)

	router.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))

}
