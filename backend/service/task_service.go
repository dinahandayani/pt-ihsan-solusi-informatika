package service

import (
	"net/http"
	"ptihsan/database"
	"ptihsan/logs"
	"ptihsan/model"
	"strconv"
)

func CreateTask(r *http.Request, task model.Task) (*model.Task, error) {
	task.Completed = false // pastikan selalu false
	if err := database.DB.Create(&task).Error; err != nil {
		logs.LogError(r, "Gagal membuat task di database: "+err.Error(), http.StatusInternalServerError, 0)
		return nil, err
	}
	logs.LogInfo(r, "Task berhasil dibuat di database, ID: "+strconv.Itoa(int(task.ID)), http.StatusCreated, 0)
	return &task, nil
}

func UpdateTask(r *http.Request, id uint, title string) (*model.Task, error) {
	var existingTask model.Task
	if err := database.DB.First(&existingTask, id).Error; err != nil {
		logs.LogError(r, "Task tidak ditemukan untuk update, ID: "+strconv.Itoa(int(id)), http.StatusNotFound, 0)
		return nil, err
	}

	existingTask.Title = title // hanya update title

	if err := database.DB.Save(&existingTask).Error; err != nil {
		logs.LogError(r, "Gagal memperbarui task ID "+strconv.Itoa(int(id))+": "+err.Error(), http.StatusInternalServerError, 0)
		return nil, err
	}

	logs.LogInfo(r, "Berhasil memperbarui task ID: "+strconv.Itoa(int(existingTask.ID)), http.StatusOK, 0)
	return &existingTask, nil
}

func DeleteTask(r *http.Request, id uint) error {
	if err := database.DB.Delete(&model.Task{}, id).Error; err != nil {
		logs.LogError(r, "Gagal menghapus task ID "+strconv.Itoa(int(id))+": "+err.Error(), http.StatusInternalServerError, 0)
		return err
	}
	logs.LogInfo(r, "Berhasil menghapus task ID: "+strconv.Itoa(int(id)), http.StatusNoContent, 0)
	return nil
}

func GetOngoingTasks(r *http.Request) ([]model.Task, error) {
	var tasks []model.Task
	if err := database.DB.
		Where("completed = ?", false).
		Order("created_at ASC").
		Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetCompletedTasks(r *http.Request) ([]model.Task, error) {
	var tasks []model.Task
	if err := database.DB.
		Where("completed = ?", true).
		Order("updated_at DESC").
		Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func MarkTaskAsCompleted(r *http.Request, id uint) error {
	var task model.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		return err
	}
	task.Completed = true
	if err := database.DB.Save(&task).Error; err != nil {
		return err
	}
	return nil
}

func GetTaskByID(r *http.Request, id uint) (*model.Task, error) {
	var task model.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		logs.LogError(r, "Task tidak ditemukan, ID: "+strconv.Itoa(int(id)), http.StatusNotFound, 0)
		return nil, err
	}
	return &task, nil
}
