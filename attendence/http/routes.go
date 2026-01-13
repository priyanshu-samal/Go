package http

import "net/http"

func NewRouter() http.Handler{
     mux:=http.NewServeMux()
	 
	 //frontend
	 mux.HandleFunc("/signup",SignupPage)
	 mux.HandleFunc("/login",LoginPage)
	 
	 //dashboard
	 mux.HandleFunc("/teacher",TeachDash)
	 mux.HandleFunc("/student",StuDash)

	 //api
	 mux.HandleFunc("/api/signup",SignupHandler)
	 mux.HandleFunc("/api/login", LoginHandler)
	 

}
