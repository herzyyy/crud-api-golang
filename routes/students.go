package routes

import (
	"api/controllers"
	"net/http"
)

func SetupStudentRoutes() {
    http.HandleFunc("GET /api/students/", controllers.GetStudents)
    http.HandleFunc("POST /api/students/", controllers.PostStudents)
    http.HandleFunc("PUT /api/students/", controllers.PutStudents)
    http.HandleFunc("DELETE /api/students/", controllers.DeleteStudents)
	
	// http.HandleFunc("/api/students/", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		controllers.GetStudents(w, r)
	// 	case http.MethodPost:
	// 		controllers.PostStudents(w, r)
	// 	case http.MethodPut:
	// 		controllers.PutStudents(w, r)
	// 	case http.MethodDelete:
	// 		controllers.DeleteStudents(w, r)
	// 	default:
	// 		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	// 	}
	// })
}
