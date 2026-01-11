package main

import (
	"context"
	
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

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

    slog.Info("Server started %s", slog.String("adress", cfg.Addr))
    

    done:= make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

    go func (){
		err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal("Failed to start server")
	}
	}()

    <-done
	slog.Info("Shutting down server")
    ctx, cancel:=	context.WithTimeout(context.Background(), 5*time.Second)
	
    defer cancel()

	err:=server.Shutdown(ctx)

	if err !=nil {
		slog.Error("Failed to Shutdown", slog.String("error", err.Error()))
	}
   
   slog.Info("Server shutdown successfully")
	
}