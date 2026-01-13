package app

import "net/http"

func NewServer() *http.Server{

	return &http.Server{
		Addr: "8080",
        Handler: router,
	}
}