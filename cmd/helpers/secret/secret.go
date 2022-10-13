// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package secret

import (
	"encoding/base64"
)

func Encrypt(plainText string) (cypherText string) {
	cypherText = encrypt(plainText)
	cypherText = base64.StdEncoding.EncodeToString([]byte(cypherText))

	return
}

func Decrypt(cypherText string) (secret string) {
	bytes, err := base64.StdEncoding.DecodeString(cypherText)
	checkErr(err)
	secret = decrypt(string(bytes))

	return
}
