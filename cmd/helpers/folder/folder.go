package folder

import (
	"fmt"
	"os"
)

func Initialize(handler func(err error)) {
	if handler == nil {
		panic("Please provide an error handler")
	}
	errorHandlerCallback = handler
}

func MkdirAll(folder string) {
	if folder == "" {
		return
	}
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("Unable to create folder '%v'. %v", folder, err))
		}
	}
}
