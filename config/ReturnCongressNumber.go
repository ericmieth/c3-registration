package config

import (
	"encoding/json"
	"log"
	"os"
)

func ReturnCongressNumber() string {

	var c Config

	file, err := os.Open("config.json")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	if err != nil {
		log.Fatal(err)
	}

	err = decoder.Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c.CongressNumber

}
