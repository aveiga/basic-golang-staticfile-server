package customerrors

import "fmt"

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
}

func (r *RestError) Error() string {
	return fmt.Sprintf("%v; Status: %v; Code: %v", r.Message, r.Status, r.Code)
}
