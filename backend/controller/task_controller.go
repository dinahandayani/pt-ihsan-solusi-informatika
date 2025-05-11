package controller

import (
	"encoding/json"
	"net/http"
	"ptihsan/logs"
	"ptihsan/model"
	"ptihsan/service"
	"ptihsan/utils"
	"strconv"
	"time"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Ambil input dari body request
	var input struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logs.LogError(r, "Bad Request: "+err.Error(), http.StatusBadRequest, time.Since(start))
		utils.RespondError(w, http.StatusBadRequest, "Permintaan tidak valid")
		return
	}

	// Buat task baru
	task := model.Task{
		Title:     input.Title,
		Completed: false,
	}

	newTask, err := service.CreateTask(r, task)
	if err != nil {
		logs.LogError(r, "Error creating task: "+err.Error(), http.StatusInternalServerError, time.Since(start))
		utils.RespondError(w, http.StatusInternalServerError, "Gagal membuat task")
		return
	}

	logs.LogInfo(r, "Task created successfully, ID: "+strconv.Itoa(int(newTask.ID)), http.StatusCreated, time.Since(start))
	utils.RespondSuccess(w, http.StatusCreated, "Task berhasil dibuat", newTask)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Ambil ID dari query
	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		logs.LogError(r, "Invalid Task ID", http.StatusBadRequest, time.Since(start))
		utils.RespondError(w, http.StatusBadRequest, "Task ID tidak valid")
		return
	}

	// Ambil input title dari body
	var input struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logs.LogError(r, "Bad Request", http.StatusBadRequest, time.Since(start))
		utils.RespondError(w, http.StatusBadRequest, "Permintaan tidak valid")
		return
	}

	// Update task melalui service
	updatedTask, err := service.UpdateTask(r, uint(id), input.Title)
	if err != nil {
		logs.LogError(r, "Error updating task", http.StatusInternalServerError, time.Since(start))
		utils.RespondError(w, http.StatusInternalServerError, "Gagal mengupdate task")
		return
	}

	logs.LogInfo(r, "Task updated successfully", http.StatusOK, time.Since(start))
	utils.RespondSuccess(w, http.StatusOK, "Task berhasil diupdate", updatedTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Ambil ID dari query
	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		logs.LogError(r, "Invalid Task ID", http.StatusBadRequest, time.Since(start))
		utils.RespondError(w, http.StatusBadRequest, "Task ID tidak valid")
		return
	}

	// Hapus task via service
	if err := service.DeleteTask(r, uint(id)); err != nil {
		logs.LogError(r, "Gagal menghapus task", http.StatusInternalServerError, time.Since(start))
		utils.RespondError(w, http.StatusInternalServerError, "Gagal menghapus task")
		return
	}

	logs.LogInfo(r, "Task deleted successfully", http.StatusOK, time.Since(start))
	utils.RespondSuccess(w, http.StatusOK, "Task berhasil dihapus", nil)
}

func GetOngoingTasks(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	tasks, err := service.GetOngoingTasks(r)
	if err != nil || len(tasks) == 0 {
		logs.LogError(r, "Data task ongoing tidak ditemukan", http.StatusNotFound, time.Since(start))
		utils.RespondError(w, http.StatusNotFound, "Data task ongoing tidak ditemukan")
		return
	}

	logs.LogInfo(r, "Fetched ongoing tasks", http.StatusOK, time.Since(start))
	utils.RespondSuccess(w, http.StatusOK, "Data task ongoing ditemukan", tasks)
}

func GetCompletedTasks(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	tasks, err := service.GetCompletedTasks(r)
	if err != nil {
		logs.LogError(r, "Error fetching completed tasks: "+err.Error(), http.StatusInternalServerError, time.Since(start))
		utils.RespondError(w, http.StatusInternalServerError, "Terjadi kesalahan saat mengambil data task completed")
		return
	}

	if len(tasks) == 0 {
		logs.LogError(r, "Data task completed tidak ditemukan", http.StatusNotFound, time.Since(start))
		utils.RespondError(w, http.StatusNotFound, "Data task completed tidak ditemukan")
		return
	}

	logs.LogInfo(r, "Fetched completed tasks", http.StatusOK, time.Since(start))
	utils.RespondSuccess(w, http.StatusOK, "Data task completed ditemukan", tasks)
}

func SignToCompleted(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Ambil ID dari query
	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		logs.LogError(r, "Invalid Task ID", http.StatusBadRequest, time.Since(start))
		utils.RespondError(w, http.StatusBadRequest, "Task ID tidak valid")
		return
	}

	// Tandai task sebagai completed
	if err := service.MarkTaskAsCompleted(r, uint(id)); err != nil {
		logs.LogError(r, "Gagal menandai task sebagai selesai", http.StatusInternalServerError, time.Since(start))
		utils.RespondError(w, http.StatusInternalServerError, "Gagal menandai task sebagai selesai")
		return
	}

	logs.LogInfo(r, "Task marked as completed", http.StatusOK, time.Since(start))
	utils.RespondSuccess(w, http.StatusOK, "Task berhasil ditandai sebagai selesai", nil)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Ambil ID dari query parameter
	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		logs.LogError(r, "Invalid Task ID", http.StatusBadRequest, time.Since(start))
		utils.RespondError(w, http.StatusBadRequest, "Task ID tidak valid")
		return
	}

	// Ambil task dari service berdasarkan ID
	task, err := service.GetTaskByID(r, uint(id))
	if err != nil {
		logs.LogError(r, "Task not found", http.StatusNotFound, time.Since(start))
		utils.RespondError(w, http.StatusNotFound, "Task tidak ditemukan")
		return
	}

	logs.LogInfo(r, "Fetched task by ID", http.StatusOK, time.Since(start))
	utils.RespondSuccess(w, http.StatusOK, "Task ditemukan", task)
}
