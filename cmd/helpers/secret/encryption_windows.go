package secret

import (
	"github.com/billgraziano/dpapi"
)

func encrypt(plainText string) (cypherText string) {
	var err error

	cypherText, err = dpapi.Encrypt(plainText)
	checkErr(err)

	return
}

func decrypt(cypherText string) (secret string) {
	var err error

	secret, err = dpapi.Decrypt(cypherText)
	checkErr(err)

	return
}
