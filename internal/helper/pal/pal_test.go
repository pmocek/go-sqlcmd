package pal

import (
	"errors"
	"testing"
)

func TestFilenameInUserHomeDotDirectory(t *testing.T) {
	FilenameInUserHomeDotDirectory(".foo", "bar")
}

func TestCheckErr(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	checkErr(errors.New("test"))
}
