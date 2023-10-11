package config

import (
	"os"
	"os/user"
	"path/filepath"
)

var (
	rootDir string
)

// RootDir returns absolute path of config directory
func RootDir() string {
	return rootDir
}

// Parse config variables
func Parse() error {
	var dir string

	usr, err := user.Current()
	if err != nil {
		dir = os.TempDir()
	} else {
		dir = usr.HomeDir
	}

	rootDir = filepath.Join(dir, ".fconsole")
	return nil
}
