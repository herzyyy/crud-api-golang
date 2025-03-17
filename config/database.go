package config

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
    // Konfigurasi koneksi database
    dsn := "root:@tcp(localhost:3306)/learn_"
    
    // Membuat koneksi ke database
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(err.Error())
    }
    
    // Test koneksi
    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }
    
    return db
} 