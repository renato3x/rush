package server

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("Handling connection from %s", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	cache := make(map[string]string)
	for {
		conn.Write([]byte("-> "))
		cmd, _ := reader.ReadString('\n')
		normalizedCmd := strings.ToLower(strings.TrimSpace(cmd))

		switch {
		case normalizedCmd == "ping":
			conn.Write([]byte("PONG\n"))
		case strings.HasPrefix(normalizedCmd, "set"):
			splitCommand := strings.Split(normalizedCmd, " ")
			if len(splitCommand) < 3 {
				conn.Write([]byte("INVALID COMMAND: " + normalizedCmd + "\n"))
				continue
			}

			key := splitCommand[1]
			value := splitCommand[2]
			cache[key] = value
		case strings.HasPrefix(normalizedCmd, "get"):
			splitCommand := strings.Split(normalizedCmd, " ")
			if len(splitCommand) < 2 {
				conn.Write([]byte("INVALID COMMAND: " + normalizedCmd + "\n"))
			}
			key := splitCommand[1]
			value, ok := cache[key]
			if !ok {
				conn.Write([]byte("INVALID COMMAND: KEY " + key + " not found \n"))
			}

			conn.Write([]byte(value + "\n"))
		case strings.HasPrefix(normalizedCmd, "del"):
			splitCommand := strings.Split(normalizedCmd, " ")
			if len(splitCommand) < 2 {
				conn.Write([]byte("INVALID COMMAND: " + cmd + "\n"))
			}
			key := splitCommand[1]
			delete(cache, key)
			conn.Write([]byte("OK\n"))
		default:
			conn.Write([]byte("INVALID COMMAND: " + cmd + "\n"))
		}
	}
}
