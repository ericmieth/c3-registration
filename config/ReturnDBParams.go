package config

import (
	"encoding/json"
	"log"
	"os"
)

func ReturnDBParams() map[string]string {

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

	configMap := make(map[string]string)
	configMap["DBHost"] = c.DBHost
	configMap["DBPort"] = c.DBPort
	configMap["DBName"] = c.DBName
	configMap["DBUser"] = c.DBUser
	configMap["DBPass"] = c.DBPass

	return configMap

}
