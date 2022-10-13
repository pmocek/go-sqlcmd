// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package formatter

import (
	"encoding/json"
	"fmt"
)

type Json struct {
	Base
}

func (f *Json) Serialize(in interface{}) {
	bytes, err := json.MarshalIndent(in, "", "    ")
	f.Base.CheckErr(err)
	fmt.Println(string(bytes))
	_, err = fmt.Println(string(bytes))
	f.Base.CheckErr(err)
}
