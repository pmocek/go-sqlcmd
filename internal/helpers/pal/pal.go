package pal

import (
	"os"
	"path/filepath"
)

func FilenameInUserHomeDotDirectory(dotDirectory string, filename string) string {
	home, err := os.UserHomeDir()
	checkErr(err)
	return filepath.Join(home, dotDirectory, filename)
}
