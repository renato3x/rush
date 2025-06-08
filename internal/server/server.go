package server

import (
	"log"
	"net"
	"strconv"
)

func Run(port int) {
	strPort := strconv.Itoa(port)
	listener, err := net.Listen("tcp", ":"+strPort)
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()
	log.Printf("rush server is open on port %s\n", strPort)
	log.Printf("Accepting connections")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s\n", err)
			messageError(conn, "Connection not accepted")
			continue
		}

		go handleConnection(conn)
	}
}
