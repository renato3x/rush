package server

import (
	"fmt"
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
	log.Printf("rush running on port %s\n", strPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: " + err.Error())
			continue
		}

		go handleConnection(conn)
	}
}
