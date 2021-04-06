package httpserver

import (
	"net/http"
	"strings"
)

func AccessControlAllowMethods() string {
	var method = []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPost,
		http.MethodOptions,
	}
	return strings.Join(method, ",")
}
