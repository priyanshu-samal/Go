package main

import (
	"fmt"
	"net/http"
	"github.com/priyanshu-samal/Intro/internal/routes"
)

func main() {
	router:=routes.NewRouter()

	port := ":8080"
	addr:= fmt.Sprintf("Listening on http://localhost%s", port)
	fmt.Printf(`Server started. %s`, addr)
	if err := http.ListenAndServe(port, router); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}

}