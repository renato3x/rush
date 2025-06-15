package server

import (
	"fmt"
	"net"
	"rush/internal/persistence"
	"strings"
)

func process(fullCommand string) (string, []string, error) {
	trimmedCommand := strings.TrimSpace(fullCommand)
	splitCommand := strings.Split(trimmedCommand, " ")

	if len(splitCommand) == 1 && splitCommand[0] == "" {
		err := fmt.Errorf("no command given")
		return "", nil, err
	}

	command := splitCommand[0]
	var args []string
	if len(splitCommand) > 1 {
		args = splitCommand[1:]
	}

	return command, args, nil
}

func cmd(conn net.Conn, fullCommand string) {
	command, args, err := process(fullCommand)

	if err != nil {
		return
	}

	if command == "ping" {
		message(conn, "PONG")
		return
	}

	if command == "get" {
		if len(args) != 1 {
			messageError(conn, "ERR usage: GET <key>")
			return
		}

		value := get(args[0])
		message(conn, value)
		return
	}

	if command == "set" {
		if len(args) != 2 {
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
