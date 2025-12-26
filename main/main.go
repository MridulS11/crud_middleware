package main

import (
	"crud_api/configs"
	"crud_api/internals/handlers"
	"crud_api/internals/middleware"
	"fmt"
	"log"
	"net/http"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("POST /books", handlers.AddBook)
	mux.HandleFunc("GET /books", handlers.GetBook)
	mux.HandleFunc("GET /books/{id}", handlers.GetBookById)
	mux.HandleFunc("PUT /books/{id}", handlers.UpdateBook)
	mux.HandleFunc("PATCH /books/{id}", handlers.UpdateBookPart)
	mux.HandleFunc("DELETE /books/{id}", handlers.DeleteBook)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Client!")
	})
	middle := middleware.JsonContentType(middleware.SecLayer(mux))
	if err := http.ListenAndServe(":8080", middle); err != nil{
		log.Println(configs.ErrCode,err)
		return
	}
}