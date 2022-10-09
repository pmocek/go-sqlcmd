package net

import (
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output"
	"net"
	"strconv"
	"time"
)

func IsLocalPortAvailable(port int) (portAvailable bool) {
	timeout := time.Second
	output.Tracef(
		"Checking if local port %d is available using DialTimeout(tcp, %v, timeout: %d)",
		port,
		net.JoinHostPort("localhost", strconv.Itoa(port)),
		timeout,
	)
	conn, err := net.DialTimeout(
		"tcp",
		net.JoinHostPort("localhost", strconv.Itoa(port)),
		timeout,
	)
	if err != nil {
		output.Tracef(
			"Expected connecting error '%v' connecting to local port %d, therefore port is available)",
			err,
			port,
		)
		portAvailable = true
	}
	if conn != nil {
		conn.Close()
		output.Tracef("Port '%d' is not available. Opened", port, net.JoinHostPort("localhost", strconv.Itoa(port)))
	} else {
		output.Tracef("Local port '%d' is available", port)
	}

	return
}
