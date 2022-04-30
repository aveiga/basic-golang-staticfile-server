package uaa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type TokenIntrospectionResult struct {
	Active bool   `json:"active"`
	Scope  string `json:"scope"`
}

func introspect(token string) (TokenIntrospectionResult, error) {
	introspectionURL := strings.Split(os.Getenv("OAUTH2_TOKEN_INTROSPECTION_URL"), "/auth")

	apiUrl := introspectionURL[0]
	data := url.Values{}
	data.Set("token", token)
	data.Set("client_id", os.Getenv("OAUTH2_CLIENT"))
	data.Set("client_secret", os.Getenv("OAUTH2_CLIENT_SECRET"))

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = fmt.Sprintf("/auth%v", introspectionURL[1])
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

func getRawToken(c *gin.Context) (string, error) {
	token := c.Request.Header.Get("Authorization")
	if strings.HasPrefix(token, "Bearer ") {
		var rawToken = strings.TrimPrefix(token, "Bearer ")
		return rawToken, nil
	}

	return "", fmt.Errorf("No Bearer Token")
}

func isElementInSlice(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawToken, err := getRawToken(c)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		introspectionResult, err := introspect(rawToken)
		if err != nil || introspectionResult.Active == false {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			c.Next()
		}

	}
}

func HasScope(scope string) gin.HandlerFunc {
	return func(c *gin.Context) {
		rawToken, err := getRawToken(c)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		introspectionResult, err := introspect(rawToken)
		if err != nil || introspectionResult.Active == false || !isElementInSlice(strings.Split(introspectionResult.Scope, " "), scope) {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			c.Next()
		}
	}
}
