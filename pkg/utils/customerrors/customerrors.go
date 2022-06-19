package customerrors

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	CategoryBusiness      string = "Business Error"
	CategoryAccessControl string = "Access Control"
	CategoryGeneral       string = "General"
)

type RestError struct {
	Status        int      `json:"status"`
	ErrorMessages []string `json:"errorMessage"`
	Category      string   `json:"category"`
	Username      string   `json:"username"`
}

func (r *RestError) Error() string {
	return fmt.Sprintf("Error Messages: %v; Status: %v; Category: %s, Username: %s", r.ErrorMessages, r.Status, r.Category, r.Username)
}

func ErrorHandler(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, ginErr := range c.Errors {
			var unwrap *RestError
			if errors.As(ginErr, &unwrap) {
				logger.Error(unwrap)
			} else {
				logger.Error(ginErr.Error())
			}
		}

		if len(c.Errors) > 0 {
			var unwrap *RestError
			if errors.As(c.Errors[0], &unwrap) {
				// status -1 doesn't overwrite existing status code
				c.JSON(-1, unwrap)
			} else {
				c.JSON(-1, c.Errors[0])
			}
		}
	}
}

func NewErrorMessageList(errorMessages ...string) []string {
	return errorMessages
}
