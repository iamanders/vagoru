package database

import (
	"database/sql"
	"flag"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestDbPath(t *testing.T) {
	var p = getDbPath() // Should return :memory: in testing
	if p != ":memory:" {
		t.Errorf("%s should be equal to :memory:", p)
	}

	// Reset flags - https://stackoverflow.com/a/57972717
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	p = getDbPath() // Should not be :memory:
	if p == ":memory:" {
		t.Errorf("%s should not be equal to :memory:", p)
	}
	if !strings.HasSuffix(p, "/.vagoru.db") {
		t.Errorf("%s should end with /.vagoru.db", p)
	}
}

func TestOpenDb(t *testing.T) {
	var dbHandle, _ = OpenDb()
	if reflect.TypeOf(dbHandle) != reflect.TypeOf(&sql.DB{}) {
		t.Error("OpenDB didn't return *sql.DB")
	}
}

// func TestIsDbInitiated() {}

// func TestInitDb() {}

// func TestCloseDb() {}

// func TestGetDb() {}
