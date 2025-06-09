package server

import (
	"bufio"
	"github.com/common-nighthawk/go-figure"
	"log"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr()
	log.Printf("Handling connection from %s", addr)

	reader := bufio.NewReader(conn)
	ascii := figure.NewFigure("RUSH", "", true).String()
	message(conn, ascii+"\n")

	for {
		conn.Write([]byte("- "))
		fullCommand, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Client from %s has disconnected", addr)
			return
		}

		splitCommand := strings.Fields(fullCommand)

		args := make([]string, len(splitCommand))
		if len(splitCommand) > 1 {
			args = splitCommand[1:]
		}

		cmd(conn, splitCommand[0], args...)
	}
}
