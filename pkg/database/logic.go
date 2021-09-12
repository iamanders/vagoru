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

func StartTracking() {

}
