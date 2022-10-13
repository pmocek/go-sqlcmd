// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package file

import (
	"github.com/microsoft/go-sqlcmd/cmd/helpers/folder"
	"os"
	"path/filepath"
)

func CreateEmptyFileIfNotExists(filename string) {
	if filename == "" {
		panic("filename must not be empty")
	}

	d, _ := filepath.Split(filename)
	if !Exists(d) {
		folder.MkdirAll(d)
	}
	if !Exists(filename) {
		handle, err := os.Create(filename)
		checkErr(err)
		defer handle.Close()
	}
}

func Exists(filename string) (exists bool) {
	if filename == "" {
		panic("filename must not be empty")
	}

	if _, err := os.Stat(filename); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}
