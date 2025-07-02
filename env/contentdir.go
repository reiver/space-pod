package env

import (
	"os"
)

var ContentDir string = contentDir()

// contentDir returns the path to the content-dir.
// I.e., where the content is stored as a content-addressable storage (CAS).
//
// It defaults to "./contents".
//
// But that can be overridden by a value set in the "CONTENT_DIR" environment variable.
func contentDir() string {
	value := os.Getenv("CONTENT_DIR")
	if "" == value {
		value = "./contents"
	}

	return value
}
