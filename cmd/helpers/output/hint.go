// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package output

var hintCallback func(hints []string)

func displayHints(hints []string) {
	if hintCallback == nil {
		panic("hintCallback not initialized")
	}

	hintCallback(hints)
}
