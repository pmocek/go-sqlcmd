package secret

func encrypt(plainText string) (cypherText string) {

	//BUG(stuartpa): Encryption not yet implemented on MacOS, will use the KeyChain
	cypherText = plainText

	return
}

func decrypt(cypherText string) (secret string) {
	secret = cypherText

	//BUG(stuartpa): Encryption not yet implemented on MacOS, will use the KeyChain
	return
}
