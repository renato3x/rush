package server

import (
	"fmt"
	"net"
	"rush/internal/persistence"
	"strconv"
	"strings"
	"time"
)

func process(fullCommand string) (string, []string, error) {
	trimmedCommand := strings.TrimSpace(fullCommand)
	splitCommand := strings.Split(trimmedCommand, " ")

	if len(splitCommand) == 1 && splitCommand[0] == "" {
		err := fmt.Errorf("no command given")
		return "", nil, err
	}

	command := strings.ToLower(splitCommand[0])
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
			messageError(conn, "Invalid usage: GET <key>")
			return
		}

		value := get(args[0])
		message(conn, value)
		return
	}

	if command == "set" {
		if len(args) < 2 {
			messageError(conn, "Invalid usage: SET <key> <value>")
			return
		}

		key, value := args[0], args[1]
		ttl := -1

		if len(args) > 2 {
			ttl, err = strconv.Atoi(args[2])
			if err != nil {
				messageError(conn, "Invalid usage: Invalid TTL "+args[2])
				return
			}
		}

		err := set(key, value, ttl)
		if err != nil {
			messageError(conn, "Cannot persist KEY "+key)
			return
		}

		message(conn, "OK")
		return
	}

	if command == "del" {
		if len(args) != 1 {
			messageError(conn, "Invalid usage: DEL <key>")
			return
		}

		key := args[0]
		del(key)
		message(conn, "OK")
		return
	}

	if command == "size" {
		fullSize := size()
		message(conn, strconv.Itoa(fullSize))
		return
	}

	messageError(conn, "Invalid command \"%s\"", command)
}

func set(key, value string, ttl int) error {
	persistence.Data[key] = value
	err := persistence.Save()
	if err != nil {
		delete(persistence.Data, key)
		return err
	}

	if ttl > -1 {
		go func() {
			time.Sleep(time.Duration(ttl) * time.Second)
			del(key)
		}()
	}

	return nil
}

func del(key string) {
	delete(persistence.Data, key)
	persistence.Save()
}

func get(key string) string {
	value, found := persistence.Data[key]
	if !found {
		return "nil"
	}
	return value
}

func size() int {
	return len(persistence.Data)
}
