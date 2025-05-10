package router

import (
	"net/http"
	"ptihsan/controller"
)

func SetupRoutes() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controller.GetTasks(w, r)
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
}
