// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package secret

import (
	"encoding/base64"
)

func Encrypt(plainText string) (cipherText string) {
	cipherText = encrypt(plainText)
	cipherText = base64.StdEncoding.EncodeToString([]byte(cipherText))

	return
}

func Decrypt(cipherText string) (secret string) {
	bytes, err := base64.StdEncoding.DecodeString(cipherText)
	checkErr(err)
	secret = decrypt(string(bytes))

	return
}
