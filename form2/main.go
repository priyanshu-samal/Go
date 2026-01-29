package main

import "net/http"

func submit(w http.ResponseWriter, r* http.Request){
       if r.Method!=http.MethodGet{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	   }
	   err:=r.ParseForm()
	   if err!=nil{
		http.Error(w,"nothing found",http.StatusBadGateway)
		return
	   }
	   name:=r.FormValue("name")
	   email:=r.FormValue("email")
	   
}

func main() {

    http.HandleFunc("/",submit)

	http.ListenAndServe(":8080",nil)
}