package main

import (
	"fmt"
	"net/http"
	"task-manager/internal/handlers"
)

func main() {
	fmt.Println("Server starting on port 8000...")
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.CreateTaskHandler(w, r)
		} else if r.Method == http.MethodGet {
			handlers.GetTasksHandler(w, r)
		} else {
			http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			handlers.DeleteTasksHandler(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	//Listen on port 8000
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
