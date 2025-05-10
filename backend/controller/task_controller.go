package controller

import (
	"encoding/json"
	"net/http"
	"ptihsan/logs"
	"ptihsan/model"
	"ptihsan/service"
	"strconv"
	"time"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	start := time.Now() // Track start time

	var task model.Task
	// Parse request body
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		// Log error dengan parameter yang sesuai
		logs.LogError(r, "Bad Request: "+err.Error(), http.StatusBadRequest, time.Since(start))
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Panggil service untuk membuat task baru, sekarang dengan request r
	newTask, err := service.CreateTask(r, task) // Menambahkan r sebagai parameter
	if err != nil {
		logs.LogError(r, "Error creating task: "+err.Error(), http.StatusInternalServerError, time.Since(start))
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}

	// Kirim response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

	// Log sukses
	logs.LogInfo(r, "Task created successfully, ID: "+strconv.Itoa(int(newTask.ID)), http.StatusCreated, time.Since(start))
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	start := time.Now() // Track start time

	// Panggil service untuk mengambil semua tasks, sekarang dengan request r
	tasks, err := service.GetTasks(r) // Menambahkan r sebagai parameter
	if err != nil {
		logs.LogError(r, "Error fetching tasks: "+err.Error(), http.StatusInternalServerError, time.Since(start))
		http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
		return
	}

	// Kirim response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

	// Log sukses
	logs.LogInfo(r, "Fetched all tasks, total: "+strconv.Itoa(len(tasks)), http.StatusOK, time.Since(start))
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	start := time.Now() // Track start time

	// Mengambil ID dari URL query parameter
	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		logs.LogError(r, "Invalid Task ID", http.StatusBadRequest, time.Since(start))
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	var task model.Task
	// Parse request body
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		logs.LogError(r, "Bad Request", http.StatusBadRequest, time.Since(start))
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Panggil service untuk memperbarui task, sekarang dengan request r
	updatedTask, err := service.UpdateTask(r, uint(id), task) // Menambahkan r sebagai parameter
	if err != nil {
		logs.LogError(r, "Error updating task", http.StatusInternalServerError, time.Since(start))
		http.Error(w, "Error updating task", http.StatusInternalServerError)
		return
	}

	// Kirim response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)

	// Log sukses
	logs.LogInfo(r, "Task updated successfully", http.StatusOK, time.Since(start))
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	start := time.Now() // Track start time

	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		logs.LogError(r, "Invalid Task ID", http.StatusBadRequest, time.Since(start))
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	// Panggil service untuk menghapus task
	if err := service.DeleteTask(r, uint(id)); err != nil {
		logs.LogError(r, "Error deleting task", http.StatusInternalServerError, time.Since(start))
		http.Error(w, "Error deleting task", http.StatusInternalServerError)
		return
	}

	// Kirim response
	w.WriteHeader(http.StatusNoContent)

	// Log sukses
	logs.LogInfo(r, "Task deleted successfully", http.StatusNoContent, time.Since(start))
}
