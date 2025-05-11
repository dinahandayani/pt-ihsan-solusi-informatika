package router

import (
	"net/http"
	"ptihsan/controller"
)

func SetupRoutes() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controller.CreateTask(w, r)
		case http.MethodPut:
			controller.UpdateTask(w, r)
		case http.MethodDelete:
			controller.DeleteTask(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/tasks-ongoing", controller.GetOngoingTasks)
	http.HandleFunc("/tasks-completed", controller.GetCompletedTasks)
	http.HandleFunc("/sign-to-completed", controller.SignToCompleted)
	http.HandleFunc("/tasks-id", controller.GetTaskByID)
}
