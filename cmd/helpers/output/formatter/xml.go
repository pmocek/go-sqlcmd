// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package formatter

import (
	"encoding/xml"
	"fmt"
)

type Xml struct {
	Base
}

func (f *Xml) Serialize(in interface{}) {
	bytes, err := xml.MarshalIndent(in,"", "    ")
	f.Base.CheckErr(err)
	_, err = fmt.Println(string(bytes))
	f.Base.CheckErr(err)
}
