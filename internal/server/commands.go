package server

import (
	"net"
	"rush/internal/persistence"
	"strings"
)

func cmd(conn net.Conn, command string, args ...string) {
	normalizedCommand := strings.ToLower(command)

	if normalizedCommand == "ping" {
		message(conn, "PONG")
		return
	}

	if normalizedCommand == "get" {
		if len(args) < 1 {
			messageError(conn, "ERR usage: GET <key>")
		}

		value := get(args[0])
		message(conn, value)
		return
	}

	if normalizedCommand == "set" {
		if len(args) < 2 {
			messageError(conn, "ERR usage: SET <key> <value>")
			return
		}

		key, value := args[0], args[1]
		err := set(key, value)
		if err != nil {
			messageError(conn, "Cannot persist KEY "+key)
			return
		}

		message(conn, "OK")
		return
	}

	messageError(conn, "Invalid command \"%s\"", command)
}

func set(key, value string) error {
	persistence.Data[key] = value
	err := persistence.Save()
	if err != nil {
		delete(persistence.Data, key)
		return err
	}
	return nil
}

func get(key string) string {
	value, found := persistence.Data[key]
	if !found {
		return "nil"
	}
	return value
}
