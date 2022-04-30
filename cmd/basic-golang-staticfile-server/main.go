package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aveiga/basic-golang-staticfile-server/internal/controllers"
	"github.com/aveiga/basic-golang-staticfile-server/internal/repositories"
	"github.com/aveiga/basic-golang-staticfile-server/internal/services"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/customamqp"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/customdb"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/uaa"
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
	defer db.Close()
	messagingClient := customamqp.NewMessagingClient()

	if dbErr != nil {
		log.Fatal("Error creating database connection")
		return
	}

	guitarRepo := repositories.NewGuitarRepo(db, ctx)
	guitarService := services.NewGuitarService(guitarRepo, messagingClient)
	guitarController := controllers.NewGuitarController(guitarService)

	router := gin.Default()

	router.POST("/guitars", guitarController.CreateGuitar)
	router.GET("/guitars", uaa.Authenticated(), guitarController.GetGuitars)
	router.GET("/guitars/:id", uaa.HasScope("email"), guitarController.SearchGuitars)
	router.DELETE("/guitars/:id", guitarController.DeleteGuitar)

	// Getting an oAuth2 Token
	// var oAuthConfig = &clientcredentials.Config{
	// 	ClientID:     os.Getenv("OAUTH2_CLIENT"),
	// 	ClientSecret: os.Getenv("OAUTH2_CLIENT_SECRET"),
	// 	TokenURL:     os.Getenv("OAUTH2_TOKEN_URL"),
	// 	Scopes:       []string{"email"},
	// }
	//
	// var token, _ = oAuthConfig.Token(ctx)
	// fmt.Printf(token.AccessToken)

	// privateService := router.Group("/api/privateService")
	// Replace the zalando.OAuth2Endpoint bellow with a Keycloak OAuth2Endpoint done by me. From the code:
	// OAuth2Endpoint is similar to the definitions in golang.org/x/oauth2
	// var OAuth2Endpoint = oauth2.Endpoint{
	// 	AuthURL:  os.Getenv("OAUTH2_AUTH_URL"),
	// 	TokenURL: os.Getenv("OAUTH2_TOKEN_URL"),
	// }
	// privateService.Use(ginoauth2.Auth(zalando.ScopeAndCheck("test", "email"), OAuth2Endpoint))
	// privateService.GET("/", func(c *gin.Context) {
	// 	if v, ok := c.Get("cn"); ok {
	// 		c.JSON(200, gin.H{"message": fmt.Sprintf("Hello from private for services to %s", v)})
	// 	} else {
	// 		c.JSON(401, gin.H{"message": "Hello from private for services without cn"})
	// 	}
	// })

	router.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))

}
