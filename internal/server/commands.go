package server

import (
	"net"
	"strings"
)

func cmd(conn net.Conn, command string, args ...string) {
	normalizedCommand := strings.ToLower(command)

	switch normalizedCommand {
	case "ping":
		message(conn, "PONG")
	default:
		messageError(conn, "Invalid command \"%s\"", command)
	}
}
