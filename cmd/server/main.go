package main

import (
	"log"
	"net/http"

	"github.com/withtahmid/crud-go/internal/database"
	"github.com/withtahmid/crud-go/internal/handlers"
)

func main(){
	database.Connect()
	
	http.HandleFunc("/list", handlers.GetList)
	
	
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}