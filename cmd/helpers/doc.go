package helpers

/*
Helpers package abstracts the following using dependency injection:

 - error handling for all non-control flow scenarios
 - trace support

This enables the application code (using these helpers) to not have to do any
error checking (except when it wants to affect application flow)

Do and Do Not:
 - Do verify parameter values and panic if the helper function cannot succeed to
   catch coding errors (do not panic for user input errors)
 - Do not output (except for the output helper). Do use the injected trace
    method to output low level debugging information
 - Do not return error if client is not going use the error for control flow, call the
   injected checkErr instead, which will probably call cobra.checkErr and exit:
     e.g. Do not sprinkle application (non-helper) code with:
       err, _ := fmt.printf("Hope this works")
       if (err != nil) {
         panic("How unlikely")
       }
     Do use the injected checkErr callback and let the application decide what to do
       err, _ := printf("Hope this works)
       checkErr(err)
 - Do not have a helper package take a dependency on another helper package
   unless they are building on each other, instead inject the needed capability in the
   help initialization
     e.g. Do not have the config helper take a dependency on the secret helper, instead
          inject the methods encrypt/decrypt to config in it's initialize method, do not:

       package config

       import (
         "github.com/microsoft/go-sqlcmd/cmd/helpers/secret"
         . "github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
       )

     Do instead:

       package config

       var encryptCallback func(plainText string) (cypherText string)
       var decryptCallback func(cipherText string) (secret string)

       func Initialize(
       encryptHandler func(plainText string)(cypherText string),
       decryptHandler func(cipherText string) (secret string),

 */
