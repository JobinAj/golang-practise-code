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

type UpdateTaskRequest struct {
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
	tasks := services.GetTask()
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
	ok := services.DeleteTask(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Write([]byte("Task deleted"))
}

//UPDATE Task

func UpdateTasksHandler(w http.ResponseWriter, r *http.Request) {
	var req UpdateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid Request", http.StatusBadRequest)
		return
	}
	parts := strings.Split(r.URL.Path, "/")
	idstr := parts[len(parts)-1]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	task, ok := services.UpdateTask(id, req.Title)
	if !ok {
		http.Error(w, "Task not found", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(task)
}
