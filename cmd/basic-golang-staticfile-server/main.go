package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aveiga/basic-golang-staticfile-server/internal/controllers"
	"github.com/aveiga/basic-golang-staticfile-server/internal/repositories"
	"github.com/aveiga/basic-golang-staticfile-server/internal/services"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/customamqp"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/customdb"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/customerrors"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/customlogger"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/uaa"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	logger := customlogger.NewCustomLogger()

	ctx := context.Background()
	db, dbErr := customdb.GetDB()
	defer db.Close()
	messagingClient := customamqp.NewMessagingClient()

	if dbErr != nil {
		logger.Fatal("Error creating database connection")
		return
	}

	guitarRepo := repositories.NewGuitarRepo(db, ctx)
	guitarService := services.NewGuitarService(guitarRepo, messagingClient)
	guitarController := controllers.NewGuitarController(guitarService)

	// gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(customerrors.ErrorHandler(logger))

	router.POST("/guitars", uaa.HasScope("email"), guitarController.CreateGuitar)
	router.GET("/guitars", uaa.Authenticated(), guitarController.GetGuitars)
	router.GET("/guitars/:id", uaa.HasScope("email"), guitarController.SearchGuitars)
	router.DELETE("/guitars/:id", guitarController.DeleteGuitar)

	// messagingClient.Subscribe("guitars", "topic", "main", func(d amqp.Delivery) {
	// 	guitar, _ := serialization.Deserialize[*models.Guitar](d.Body)
	// 	fmt.Printf(guitar.Brand)
	// })

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
