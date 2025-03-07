package helper

import (
	"net/http"

	"goexample/pkg/service_errors"
)

var StatusCodeMapping = map[string]int{
	service_errors.RecordNotFound: 404,
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return value
}
