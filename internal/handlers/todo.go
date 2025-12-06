package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

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

func GetById(w http.ResponseWriter, r *http.Request){



	// id := r.URL.Query().Get("id")

	// if id == ""{
	// 	http.Error(w, "Missing id parameter", http.StatusBadRequest)
	// 	return
	// }

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[2] == ""{
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}
	id := parts[2]

	query := "SELECT id, name, complete from todos where id = $1"
	
	var todo ToDo

	err := database.DB.QueryRow(r.Context(), query, id).Scan(
		&todo.ID, &todo.Name, &todo.Complete,
	)

	if err != nil {
		http.Error(w, "Todo Not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Application-Type", "Application.json")
	json.NewEncoder(w).Encode(todo)
}

func CreateTodo(w http.ResponseWriter, r *http.Request){
	if r.Method  != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var input struct { 
		Name string `json:"name"`
	}

	err:= json.NewDecoder(r.Body).Decode(&input)
	if err != nil{
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO todos(name, complete) VALUES($1, $2) RETURNING id, name, complete"

	var todo ToDo
	err = database.DB.QueryRow(r.Context(), query, input.Name, false).Scan(
		&todo.ID, &todo.Name, &todo.Complete,
	)

	if err != nil {
		http.Error(w, "Failed to insert", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodPut {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}

	var input ToDo

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := "UPDATE todos set name = $2, complete = $3 where id = $1 RETURNING id, name, complete"
	var todo ToDo
	err = database.DB.QueryRow(r.Context(), query, input.ID, input.Name, input.Complete).Scan(
		&todo.ID, &todo.Name, &todo.Complete,
	)
	if err != nil {
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func Delete(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodDelete{
		http.Error(w, "Invalid Method", http.StatusBadRequest)
		return
	}

	

}