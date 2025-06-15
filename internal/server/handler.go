package server

import (
	"bufio"
	"github.com/common-nighthawk/go-figure"
	"log"
	"net"
)

func welcome(conn net.Conn) {
	ascii := figure.NewFigure("RUSH", "", true).String()
	message(conn, ascii+"\n")
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr()
	log.Printf("Handling connection from %s", addr)
	welcome(conn)

	reader := bufio.NewReader(conn)
	for {
		conn.Write([]byte("- "))
		command, _ := reader.ReadString('\n')
		cmd(conn, command)
	}
}
