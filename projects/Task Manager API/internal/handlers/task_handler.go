package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"task-manager/internal/services"
)

type CreateTaskRequest struct {
	Title string `json:"title"`
}

// POST task
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) { // this is the signature for go http handler
	var req CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	task := services.CreateTask(req.Title)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// GET task
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := services.GetTasks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// DELETE task
func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	idstr := parts[len(parts)-1] //we have to specify which req len should it be calculating

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	ok := services.DeleteTasks(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Write([]byte("Task deleted"))
}
