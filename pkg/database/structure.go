package database

import "errors"

// Create the initial db structure
func CreateDbStructure() error {
	createStudentTableSQL := `CREATE TABLE times (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			"starts_at" datetime,
			"stops_at" datetime,
			"project" string,
			"comment" string);
		);`

	statement, err := dbHandle.Prepare(createStudentTableSQL)
	if err != nil {
		return errors.New("table already exists")
		// log.Fatal(err.Error())
	}
	statement.Exec()

	return nil
}
