package main

import "net/http"

func home(w http.ResponseWriter, r* http.Request){
          if r.Method!=http.MethodGet{
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		  }
		  w.Write([]byte("helooooo"))
}

func main() {
	http.HandleFunc("/",home)


	http.ListenAndServe(":8080",nil)

}