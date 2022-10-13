package secret

import (
	"encoding/base64"
	"github.com/billgraziano/dpapi"
)

func encrypt(plainText string) (cypherText string) {
	var err error

	cypherText, err = dpapi.Encrypt(plainText)
	checkErr(err)

	return
}

func decrypt(cypherText string) (secret string) {
	password, err := base64.StdEncoding.DecodeString(cypherText)
	checkErr(err)

	secret, err = dpapi.Decrypt(string(password))
	checkErr(err)

	return
}
