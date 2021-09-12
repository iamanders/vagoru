package command

import (
	"errors"

	"github.com/iamanders/vagoru/pkg/database"
)

func StartTracking(projectTitle string) (int, error) {
	sql := `INSERT INTO "times" ("starts_at", "project", "comment") VALUES ($1, $2, $3);`
	id := 0
	err := database.GetDb().QueryRow(sql, "2021-01-01 10:00:00", projectTitle, "Comment").Scan(&id)
	if err != nil {
		return -1, errors.New("could not start tracking " + err.Error())
	}

	return id, nil
}
