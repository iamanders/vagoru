package database

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var dbHandle *sql.DB
var dbPath = getDbPath()

// Get path to db
// If testing environment, :memory: is returned
func getDbPath() string {
	// Early return if testing environment
	if strings.Contains(os.Args[0], ".test") {
		return ":memory:"
	}
	// Else the database is in the home dir
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return dirname + "/.vagoru.db"
}

// Check if db file is created
func IsDbInitiated() bool {
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
}

// Close db
func CloseDb() {
	if dbHandle != nil {
		dbHandle.Close()
	}
	dbHandle = nil
}

// Get the db handle from package
func GetDb() *sql.DB {
	if dbHandle == nil {
		OpenDb()
		CreateDbStructure() // TODO: Only run if needed
	}
	return dbHandle
}

// Reset db
// Used in testing when using sqlite :memory: database
func ResetDb() {
	CloseDb()
	GetDb()
}
