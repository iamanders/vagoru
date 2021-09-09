package database

import (
	"database/sql"
	"log"
	"os"
)

// var dbPath = getHomeDir() + "/.vagoru.db"
var dbPath = "./vagoru.db"
var dbHandle *sql.DB

func getHomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return dirname
}

// Check if db file is created
func IsDbInitiated() bool {
	// if _, err := os.Stat("vagoru.db"); err == nil {
	if _, err := os.Stat(dbPath); err == nil {
		return true
	}

	return false
}

// Initialize db file
func InitDb() {
	// file, err := os.Create("vagoru.db")
	file, err := os.Create(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

// Open db
func OpenDb() (*sql.DB, error) {
	var err error
	dbHandle, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return dbHandle, nil

	// defer sqliteDatabase.Close()
}

// Close db
func CloseDb() {
	dbHandle.Close()
}

// Get the db handle from package
func GetDb() *sql.DB {
	return dbHandle
}
