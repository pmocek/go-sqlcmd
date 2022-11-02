package docker

import (
	"errors"
	"testing"
)

func Test_checkErr(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	checkErr(errors.New("verify error handler"))
}
