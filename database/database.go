package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {
	// Open database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	//  Set connection pool settings (optional but recommended)
	db.SetMaxOpenConns(25) // konfigurasi maksimal open coonection di dalam database berapa (25 transaksi dalam satu waktu)
	db.SetMaxIdleConns(5)  // konfigurasi maksimal idle connection
	log.Print("Database connected successfully")
	return db, nil
}
