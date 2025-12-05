package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/withtahmid/crud-go/internal/database"
)

type ToDo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Complete bool `json:"complete"`
}

func GetList(w http.ResponseWriter, r *http.Request){
	var query string = "SELECT * FROM todos"

	rows, err := database.DB.Query(r.Context(), query)
	if(err != nil){
		http.Error(w, "DB error", 500)
		return
	}
	defer rows.Close()

	todos := []ToDo{}

	for rows.Next() {
		var todo ToDo
		rows.Scan(&todo.ID, &todo.Name, &todo.Complete)
		todos = append(todos, todo)
	}
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(todos)
}