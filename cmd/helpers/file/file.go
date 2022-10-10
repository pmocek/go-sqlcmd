package file

import (
	"github.com/microsoft/go-sqlcmd/cmd/helpers/folder"
	"os"
	"path/filepath"
)

func Initialize(handler func(err error)) {
	if handler == nil {
		panic("Please provide an error handler")
	}
	errorHandlerCallback = handler
}

func CreateEmptyIfNotExists(filename string) {
	if !Exists(filename) {
		folder.MkdirAll(filepath.Base(filename))
		handle, err := os.Create(filename)
		checkErr(err)
		defer handle.Close()
	}
}

func Exists(filename string) (exists bool) {
	if _, err := os.Stat(filename); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}
