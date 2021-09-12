package database

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

// TODO: Move to other place
type Tracking struct {
	Title    string
	Comment  string
	StartsAt time.Time
	StopsAt  time.Time
}

func CurrentTracking() (Tracking, error) {
	stmt, err := GetDb().Prepare("SELECT project FROM times WHERE stops_at IS NULL")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var project string
	err = stmt.QueryRow().Scan(&project)
	if err != nil && err == sql.ErrNoRows {
		return Tracking{}, errors.New("no active tracking")
	}
	return Tracking{Title: project}, nil
}

func StartTracking(projectTitle string) (Tracking, error) {
	// func StartTracking(projectTitle string) (int, error) {
	sql := `INSERT INTO "times" ("starts_at", "project", "comment") VALUES ($1, $2, $3);`
	id := 0
	err := GetDb().QueryRow(sql, "2021-01-01 10:00:00", projectTitle, "Comment").Scan(&id)
	if err != nil {
		return Tracking{}, errors.New("could not start tracking " + err.Error())
	}

	return Tracking{Title: projectTitle}, nil
}
