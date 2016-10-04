package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"
)

func ReturnLoginExpireMinutes() time.Duration {

	var c Config

	//file, err := os.Open("../config.json") // for unit testing
	file, err := os.Open("config.json")
	decoder := json.NewDecoder(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = decoder.Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	duration, err := strconv.Atoi(c.LoginExpireMinutes)
	if err != nil {
		log.Print(err)
	}

	return time.Duration(duration)

}
