package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/priyanshu-samal/student/internal/config"
)

func main() {
	cfg := config.MustLoad()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Welcome"))
	})

	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	go startServer(server)

	waitForShutdown(server)
}

func startServer(server *http.Server) {
	slog.Info("HTTP server starting", slog.String("addr", server.Addr))

	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("HTTP server failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func waitForShutdown(server *http.Server) {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	slog.Info("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Graceful shutdown failed", slog.String("error", err.Error()))
		return
	}

	slog.Info("Server shut down cleanly")
}
