package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/aveiga/basic-golang-staticfile-server/pkg/models"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/customerrors"
	"github.com/gin-gonic/gin"
)

type GuitarController struct {
	guitarService models.GuitarService
}

func NewGuitarController(guitarService models.GuitarService) *GuitarController {
	return &GuitarController{
		guitarService: guitarService,
	}
}

type TokenIntrospectionResult struct {
	Active bool   `json:"active"`
	Scope  string `json:"scope"`
}

func introspect(token string) (TokenIntrospectionResult, error) {
	apiUrl := os.Getenv("OAUTH2_TOKEN_INTROSPECTION_URL")
	data := url.Values{}
	data.Set("token", token)
	data.Set("client_id", os.Getenv("OAUTH2_CLIENT"))
	data.Set("client_secret", os.Getenv("OAUTH2_CLIENT_SECRET"))

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = "/auth/realms/master/protocol/openid-connect/token/introspect"
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	// r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// Log the request body
	bodyString := string(body)
	log.Print(bodyString)
	// Unmarshal result
	res := TokenIntrospectionResult{}
	err = json.Unmarshal(body, &res)

	fmt.Printf("%v %v", res.Active, res.Scope)
	return res, err
}

func verify(c *gin.Context) bool {
	status := true
	token := c.Request.Header.Get("Authorization")
	if strings.HasPrefix(token, "Bearer ") {
		var rawToken = strings.TrimPrefix(token, "Bearer ")
		res, err := introspect(rawToken)

		if err != nil || res.Active == false {
			c.String(http.StatusUnauthorized, "Unauthorized")
			status = false
		}
	} else {
		c.String(http.StatusUnauthorized, "Unauthorized")
		status = false
	}
	return status
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
	fmt.Println(guitar)
	result, saveError := gc.guitarService.CreateGuitar(&guitar)
	if saveError != nil {
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (gc *GuitarController) GetGuitars(c *gin.Context) {
	if verify(c) {

		guitars, err := gc.guitarService.GetGuitars()
		if err != nil {
			log.Fatal(err)
			return
		}

		c.JSON(http.StatusOK, guitars)

	}
}

func (gc *GuitarController) SearchGuitars(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func (gc *GuitarController) DeleteGuitar(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
