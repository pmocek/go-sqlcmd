package pal

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func FilenameInUserHomeDotDirectory(dotDirectory string, filename string) string {
	home, err := os.UserHomeDir()
	checkErr(err)

	return filepath.Join(home, dotDirectory, filename)
}

func UserName() (userName string) {
	if runtime.GOOS == "windows" {
		userName = os.Getenv("USERNAME")
	} else {
		userName = os.Getenv("USER")
	}

	if userName == "" {
		panic("Unable to get username, set env var USERNAME or USER")
	}

	return
}

func CmdLineWithEnvVars(vars []string , cmd string) string {
	var sb strings.Builder
	for _, v := range vars {
		if runtime.GOOS == "windows" {
			sb.WriteString("SET ")
			sb.WriteString(`"` + v + `"`)
		} else {
			sb.WriteString("export ")
			sb.WriteString(`'` + v + `''`)
		}
	}

	if runtime.GOOS == "windows" {
		sb.WriteString(" & ")
	} else {
		sb.WriteString("; ")
	}
	sb.WriteString(cmd)

	return sb.String()
}
