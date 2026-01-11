package main

import (
	"log"
     "fmt"
	"github.com/priyanshu-samal/student/internal/config"

	"net/http"
)
func main() {

	cfg := config.MustLoad()

	router:=http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome"))
	}) 


	server:=http.Server{
		Addr:cfg.Addr,
		Handler: router,

	}

    fmt.Printf("Server is running %s", cfg.HTTPServer.Addr)

	err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal("Failed to start server")
	}

	
}