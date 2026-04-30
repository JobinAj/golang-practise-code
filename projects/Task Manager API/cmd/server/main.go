package main

import (
	"fmt"
	"net/http"
	"task-manager/internal/handlers"
)

func main() {
	fmt.Println("server started on 8000")

	//HELTH ENDPOINT
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// GET AND POST METHODS
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.CreateTaskHandler(w, r)
		case http.MethodGet:
			handlers.GetTasksHandler(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// DELETE AND UPDATE METHODS
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			handlers.DeleteTasksHandler(w, r)
		case http.MethodPut:
			handlers.UpdateTasksHandler(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	//Listen on port 8000
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
