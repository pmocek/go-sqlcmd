// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package secret

import (
	"encoding/base64"
)

func Encrypt(plainText string, encryptPassword bool) (cipherText string) {
	if plainText == "" {
		panic("Cannot encrypt an empty string")
	}

	if encryptPassword {
		cipherText = encrypt(plainText)
	} else {
		cipherText = plainText
	}

	cipherText = base64.StdEncoding.EncodeToString([]byte(cipherText))

	return
}

func Decrypt(cipherText string, decryptPassword bool) (plainText string) {
	if cipherText == "" {
		panic("Cannot decrypt an empty string")
	}

	bytes, err := base64.StdEncoding.DecodeString(cipherText)
	checkErr(err)

	if decryptPassword {
		plainText = decrypt(string(bytes))
	} else {
		plainText = string(bytes)
	}

	return
}
