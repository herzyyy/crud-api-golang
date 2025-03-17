package main

import (
    "fmt"
    "net/http"
    "api/routes"
)

func main() {
    // Setup routes dari package routes
    routes.SetupStudentRoutes()
    
    // Informasi endpoint yang tersedia
    fmt.Println("=========================================")
    fmt.Println("Server berjalan pada port 8080...")
    fmt.Println("Endpoint yang tersedia:")
    fmt.Println("GET: http://localhost:8080/api/students")
    fmt.Println("=========================================")
    
    http.ListenAndServe(":8080", nil)
}
