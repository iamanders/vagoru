package terminal

import (
	"os"
	"testing"
)

func TestArgv(t *testing.T) {
	// Preserve old args
	originalArgs := os.Args
	defer func() {
		os.Args = originalArgs
	}()

	// TODO
}
