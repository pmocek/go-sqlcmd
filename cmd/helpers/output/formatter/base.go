// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package formatter

type Base struct {
	ErrorHandlerCallback func(err error)
}

func (f *Base) CheckErr(err error) {
		if f.ErrorHandlerCallback == nil {
			panic("errorHandlerCallback not initialized")
		}

		f.ErrorHandlerCallback(err)
}
