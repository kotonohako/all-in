package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func DbConnection() (*sql.DB, error) {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	path := os.Getenv("DB_PATH")
	port := os.Getenv("DB_PORT")
	dataBaseName := os.Getenv("DB_DATABASE")

	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, path, port, dataBaseName)

	db, err := sql.Open("mysql", dbPath)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func DbConnectionWithSqlx() (*sqlx.DB, error) {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	path := os.Getenv("DB_PATH")
	port := os.Getenv("DB_PORT")
	dataBaseName := os.Getenv("DB_DATABASE")

	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, path, port, dataBaseName)
	db, err := sqlx.Open("mysql", dbPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}
