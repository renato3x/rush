package persistence

import (
	"encoding/json"
	"os"
)

var Data = make(map[string]string)

func Load() error {
	file, err := os.Open("db.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()
	return json.NewDecoder(file).Decode(&Data)
}

func Save() error {
	file, err := os.Create("db.json")
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(Data)
}
