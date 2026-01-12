package app

import (
	"net/http"

	internalhttp "github.com/priyanshu-samal/miniauth/internal/http"
)

func NewServer() *http.Server {
	router := internalhttp.NewRouter()

	return &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}
