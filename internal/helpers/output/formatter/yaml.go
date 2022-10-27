// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package formatter

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

type Yaml struct {
	Base
}

func (f *Yaml) Serialize(in interface{}) {
	bytes, err := yaml.Marshal(in)
	f.Base.CheckErr(err)
	_, err = fmt.Println(string(bytes))
	f.Base.CheckErr(err)
}
