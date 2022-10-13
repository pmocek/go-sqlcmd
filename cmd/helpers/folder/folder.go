// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package folder

import (
	"os"
)

func MkdirAll(folder string) {
	if folder == "" {
		panic("folder must not be empty")
	}
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err := os.MkdirAll(folder, os.ModePerm)
		checkErr(err)
	}
}
