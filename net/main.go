package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func main() {

    http.HandleFunc("/",handlefunc)

	log.Fatal(http.ListenAndServe(":8080", nil))



}

func handlefunc(w http.ResponseWriter, r *http.Request){
      wc,err:=w.Write([]byte("Hello world"))
	  if err!=nil{
		slog.Error("error writing resposnse",err)
		return
	  }
	  fmt.Println("%d bytes writen",wc)
}