// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package util

import (
	"strconv"
	"strings"
)

// errorPrefix is the prefix for all sqlcmd-generated errors
const errorPrefix = "Sqlcmd: Error: "

// argumentError is related to command line switch validation not handled by kong
type argumentError struct {
	Parameter string
	Rule      string
}

func (e *argumentError) Error() string {
	return errorPrefix + e.Rule
}

// InvalidServerName indicates the SQLCMDSERVER variable has an incorrect format
var InvalidServerName = argumentError{
	Parameter: "server",
	Rule:      "server must be of the form [tcp]:server[[/instance]|[,port]]",
}

// splitServer extracts connection parameters from a server name input
func SplitServer(serverName string) (string, string, uint64, error) {
	instance := ""
	port := uint64(0)
	if strings.HasPrefix(serverName, "tcp:") {
		if len(serverName) == 4 {
			return "", "", 0, &InvalidServerName
		}
		serverName = serverName[4:]
	}
	serverNameParts := strings.Split(serverName, ",")
	if len(serverNameParts) > 2 {
		return "", "", 0, &InvalidServerName
	}
	if len(serverNameParts) == 2 {
		var err error
		port, err = strconv.ParseUint(serverNameParts[1], 10, 16)
		if err != nil {
			return "", "", 0, &InvalidServerName
		}
		serverName = serverNameParts[0]
	} else {
		serverNameParts = strings.Split(serverName, "\\")
		if len(serverNameParts) > 2 {
			return "", "", 0, &InvalidServerName
		}
		if len(serverNameParts) == 2 {
			instance = serverNameParts[1]
			serverName = serverNameParts[0]
		}
	}
	return serverName, instance, port, nil
}

// padRight appends c instances of s to builder
func PadRight(builder *strings.Builder, c int64, s string) *strings.Builder {
	var i int64
	for ; i < c; i++ {
		builder.WriteString(s)
	}
	return builder
}

// padLeft prepends c instances of s to builder
func PadLeft(builder *strings.Builder, c int64, s string) *strings.Builder {
	newBuilder := new(strings.Builder)
	newBuilder.Grow(builder.Len())
	var i int64
	for ; i < c; i++ {
		newBuilder.WriteString(s)
	}
	newBuilder.WriteString(builder.String())
	return newBuilder
}

func Contains(arr []string, s string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}
