package server

import (
	"fmt"
	"net"
)

func Run() {
	port := "5173"
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic("Error listening: " + err.Error())
	}

	defer listener.Close()
	fmt.Println("Listening on " + port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: " + err.Error())
			continue
		}

		fmt.Println("Accepting connection "+port, conn.LocalAddr().String())
	}
}
