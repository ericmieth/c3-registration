package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"
)

func ReturnLoginExpireMinutes() time.Duration {

	type config struct {
		LoginExpireMinutes string
	}

	//file, err := os.Open("../config.json") // for unit testing
	file, err := os.Open("config.json")
	decoder := json.NewDecoder(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	configuration := config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(err)
	}

	duration, err := strconv.Atoi(configuration.LoginExpireMinutes)
	if err != nil {
		log.Print(err)
	}

	return time.Duration(duration)

}
