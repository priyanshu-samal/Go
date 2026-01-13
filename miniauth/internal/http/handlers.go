package http

import (
	"net/http"

	"github.com/priyanshu-samal/miniauth/internal/auth"
)

func SignupHandler(service *auth.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		if err := service.Signup(r.Context(), email, password); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write([]byte("signup success"))
	}
}

func LoginHandler(service *auth.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		if err := service.Login(r.Context(), email, password); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		w.Write([]byte("login success"))
	}
}
