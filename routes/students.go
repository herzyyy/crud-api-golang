package routes

import (
	"api/controllers"
	"net/http"
)

func SetupStudentRoutes() {
	http.HandleFunc("/api/students/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetStudents(w, r)
		case http.MethodPost:
			controllers.PostStudents(w, r)
		case http.MethodPut:
			controllers.PutStudents(w, r)
		case http.MethodDelete:
			controllers.DeleteStudents(w, r)
		default:
			http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		}
	})
}
