package folder

import (
	"fmt"
	"os"
)

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
