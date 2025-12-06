package main

import (
	"log"
	"net/http"

	"github.com/withtahmid/crud-go/internal/database"
	"github.com/withtahmid/crud-go/internal/handlers"
	"github.com/withtahmid/crud-go/internal/middleware"
)

func main(){
	database.Connect()
	
	mux := http.NewServeMux();


	mux.HandleFunc("/todos", handlers.GetList)
	mux.HandleFunc("/todos/update", handlers.UpdateTodo)
	mux.HandleFunc("/todos/create", handlers.CreateTodo)
	mux.HandleFunc("/todos/", handlers.GetById)
	
	handler := middleware.CORS(mux) 

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", handler)
}