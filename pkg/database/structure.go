package database

import "log"

// Create the initial db structure
func CreateDbStructure() {
	createStudentTableSQL := `CREATE TABLE times (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			"starts_at" datetime,
			"stops_at" datetime,
			"project" string,
			"comment" string);
		);`

	statement, err := dbHandle.Prepare(createStudentTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}
