package secret

func encrypt(plainText string) (cipherText string) {

	//BUG(stuartpa): Encryption not yet implemented on MacOS, will use the KeyChain
	cipherText = plainText

	return
}

func decrypt(cipherText string) (secret string) {
	secret = cipherText

	//BUG(stuartpa): Encryption not yet implemented on MacOS, will use the KeyChain
	return
}
