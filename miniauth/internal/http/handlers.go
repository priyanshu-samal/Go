package http

import (
	"fmt"
	"net/http"

	"github.com/priyanshu-samal/miniauth/internal/auth"
)

var users = map[string]string{} // name -> password hash

// ---------- PAGES (GET only) ----------

func SignupPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "web/signup.html")
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "web/login.html")
}

func DashboardPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		http.ServeFile(w, r, "web/dashboard.html")
	})
}

// ---------- API (POST only) ----------

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	password := r.FormValue("password")

	if _, exists := users[name]; exists {
		fmt.Fprintln(w, "user already exists")
		return
	}

	hash, _ := auth.HashPassword(password)
	users[name] = hash

	fmt.Fprintln(w, "signup successful, now login")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	password := r.FormValue("password")

	hash, ok := users[name]
	if !ok {
		fmt.Fprintln(w, "user not found, please signup")
		return
	}

	if err := auth.CheckPassword(hash, password); err != nil {
		fmt.Fprintln(w, "invalid password")
		return
	}

	token, _ := auth.GenerateToken(name)

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
		Path:  "/",
	})

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
