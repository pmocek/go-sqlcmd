package commander

import (
	"testing"
)

func TestAbstractBase_DefineCommand(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	c := AbstractBase{}
	c.DefineCommand()
}
