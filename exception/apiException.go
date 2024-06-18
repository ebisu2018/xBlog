package exception

import (
	"fmt"
	"net/http"
)

func NewApiException(code Code, format string, a ...any) *ApiException {
	return &ApiException{
		BusinessCode: code,
		Message:      fmt.Sprintf(format, a...),
		HttpCode: http.StatusOK,
	}
}

type ApiException struct {
	BusinessCode Code
	Message      string
	HttpCode     int
}

func (e *ApiException) Error() string {
	return fmt.Sprintf("%v - %v", e.BusinessCode, e.Message)
}
