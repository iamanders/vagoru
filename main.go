package main

import (
	"fmt"
	"log"

	"github.com/iamanders/vagoru/pkg/database"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize db
	if !database.IsDbInitiated() {
		fmt.Println("Could not find db, initializing a new")
		database.InitDb()
	}

	// Open db
	_, err := database.OpenDb()
	if err != nil {
		log.Fatal("Could not open db", err)
	}
	defer database.CloseDb()

	database.CreateDbStructure()

	// TODO: Parse terminal args
}
