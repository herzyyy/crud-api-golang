package controllers

import (
	"api/config"
	"api/models"
	"encoding/json"
	"net/http"
)

func GetStudents(w http.ResponseWriter, r *http.Request) {
	db := config.ConnectDB()
	defer db.Close()

	// Query untuk mengambil semua data students
	rows, err := db.Query("SELECT id_students, nis, name, gender FROM students")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Slice untuk menampung hasil
	var students []models.Student

	// Iterasi hasil query
	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.ID, &student.NIS, &student.Name, &student.Gender)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		students = append(students, student)
	}

	// Set header response sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response JSON
	json.NewEncoder(w).Encode(students)
}

func PostStudents(w http.ResponseWriter, r *http.Request) {
	var student models.Student

	// Decode JSON dari request body
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := config.ConnectDB()
	defer db.Close()

	// Prepare statement SQL untuk insert
	stmt, err := db.Prepare("INSERT INTO students (nis, name, gender) VALUES (?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// Eksekusi statement
	_, err = stmt.Exec(student.NIS, student.Name, student.Gender)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set header response sebagai JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Kirim response JSON
	json.NewEncoder(w).Encode(student)
}

func PutStudents(w http.ResponseWriter, r *http.Request) {
	var student models.Student

	// Decode JSON dari request body
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := config.ConnectDB()
	defer db.Close()

	// Prepare statement SQL untuk update
	stmt, err := db.Prepare("UPDATE students SET nis=?, name=?, gender=? WHERE id_students=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// Eksekusi statement
	_, err = stmt.Exec(student.NIS, student.Name, student.Gender, student.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func DeleteStudents(w http.ResponseWriter, r *http.Request) {
	var student models.Student

	// Decode JSON dari request body untuk mendapatkan ID
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := config.ConnectDB()
	defer db.Close()

	// Prepare statement SQL untuk delete
	stmt, err := db.Prepare("DELETE FROM students WHERE id_students=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// Eksekusi statement
	_, err = stmt.Exec(student.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Student berhasil dihapus"})
}
