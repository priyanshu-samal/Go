package app

import (
	"log"
	"net/http"

	internalhttp "github.com/priyanshu-samal/miniauth/internal/http"
)

func Start() {
	router := internalhttp.NewRouter()

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
