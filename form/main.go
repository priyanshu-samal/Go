package main

import ("net/http"
"fmt")

func submit(w http.ResponseWriter, r* http.Request){
	if r.Method!=http.MethodPost{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	err:=r.ParseForm()
	if err!=nil{
		http.Error(w,"error parsing form",http.StatusBadRequest)
	return
	}
	name:=r.FormValue("name")
	email:=r.FormValue("email")
	thought:=r.FormValue("thought")

	if name==""||email==""||thought==""{
		http.Error(w,"all field are required",http.StatusBadRequest)
		return
	}

      fmt.Println("Name:", name)
	fmt.Println("Email:", email)
	fmt.Println("Thought:", thought)

	w.Write([]byte("Form submitted successfully"))
	
}

func main() {
    
	http.HandleFunc("/submit",submit)

	http.ListenAndServe(":8080",nil)
}