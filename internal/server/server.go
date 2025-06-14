package server

import (
	"log"
	"net"
	"rush/internal/persistence"
	"strconv"
)

func Run(port int) {
	err := persistence.Load()
	if err != nil {
		log.Fatal(err)
	}
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
