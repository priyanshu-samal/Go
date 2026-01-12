package http

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// frontend pages
	mux.HandleFunc("/signup", SignupPage)
	mux.HandleFunc("/login", LoginPage)
	mux.Handle("/dashboard", AuthMiddleware(DashboardPage()))

	// API
	mux.HandleFunc("/api/signup", SignupHandler)
	mux.HandleFunc("/api/login", LoginHandler)

	return mux
}
