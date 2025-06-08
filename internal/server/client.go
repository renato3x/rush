package server

import (
	"fmt"
	"net"
)

func message(conn net.Conn, msg string, args ...any) {
	finalMessage := fmt.Sprintf(msg+"\n", args...)
	conn.Write([]byte(finalMessage))
}

func messageError(conn net.Conn, msg string, args ...any) {
	finalMessage := fmt.Sprintf("ERR: "+msg, args...)
	message(conn, finalMessage)
}
