package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func home(w http.ResponseWriter, r* http.Request){
	time.Sleep(3 * time.Second) // simulate work
		w.Write([]byte("hello"))
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)

	server:=&http.Server{
		Addr:"8080",
		Handler:mux,
	}

	go func(){
		log.Println("server started on :8080")
		if err := server.ListenAndServe();
		err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
	}
	}()
}