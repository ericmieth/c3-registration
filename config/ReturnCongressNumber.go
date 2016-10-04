package config

import (
	"encoding/json"
	"log"
	"os"
)

func ReturnCongressNumber() string {

	type config struct {
		CongressNumber string
	}

	file, err := os.Open("config.json")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	if err != nil {
		log.Fatal(err)
	}

	configuration := config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(err)
	}

	return configuration.CongressNumber

}
