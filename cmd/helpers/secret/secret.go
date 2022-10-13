// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package secret

import (
	"encoding/base64"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output"
)

func Encrypt(plainText string) (cipherText string) {
	cipherText = encrypt(plainText)
	cipherText = base64.StdEncoding.EncodeToString([]byte(cipherText))

	return
}

func Decrypt(cypherText string) (secret string) {
	output.Trace(cypherText)
	bytes, err := base64.StdEncoding.DecodeString(cypherText)
	checkErr(err)
	secret = decrypt(string(bytes))

	return
}
