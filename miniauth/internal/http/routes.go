package http

import (
	"net/http"

	"github.com/priyanshu-samal/miniauth/internal/auth"
	"github.com/priyanshu-samal/miniauth/internal/db"
	"github.com/priyanshu-samal/miniauth/internal/repository"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	database, _ := db.NewMongo()
	userRepo := repository.NewUserRepository(database)
	authService := auth.NewService(userRepo)

	mux.HandleFunc("/signup", SignupHandler(authService))
	mux.HandleFunc("/login", LoginHandler(authService))

	return mux
}
