// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package secret

import (
	"encoding/base64"
	"github.com/billgraziano/dpapi"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Initialize(handler errorHandlerService) {
	if handler == nil {
		panic("Please provide an error handler")
	}

	errorHandlerCallback = handler
}

func Decrypt(ciperText string) (secret string) {

	// Show password encrypted (so it can be used in other tools)
	if runtime.GOOS == "windows" {
		password, err := base64.StdEncoding.DecodeString(ciperText)
		checkErr(err)

		secret, err = dpapi.Decrypt(string(password))
		checkErr(err)
	} else {
		// TODO: MacOS (keychain) and Linux (not sure?)
		decoded, err := base64.StdEncoding.DecodeString(ciperText)
		checkErr(err)

		secret = string(decoded)
	}
	return
}

func Encrypt(plainText string) (cyperText string) {
	var err error

	if runtime.GOOS == "windows" {
		cyperText, err = dpapi.Encrypt(plainText)
		checkErr(err)
	} else {
		// TODO: MacOS (keychain) and Linux (not sure?)
		cyperText = plainText
	}

	return base64.StdEncoding.EncodeToString([]byte(cyperText))
}

// https://golangbyexample.com/generate-random-password-golang/
var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func Generate(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
