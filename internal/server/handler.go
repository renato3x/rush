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
	log.Printf("Handling connection from %s", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	ascii := figure.NewFigure("RUSH", "", true).String()
	message(conn, ascii+"\n")

	for {
		conn.Write([]byte("- "))
		fullCommand, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Client from %s has disconnected", conn.RemoteAddr())
			return
		}

		splitCommand := strings.Fields(fullCommand)
		cmd(conn, splitCommand[0], splitCommand[1:]...)
	}
}
