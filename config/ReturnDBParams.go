package config

import (
	"encoding/json"
	"log"
	"os"
)

func ReturnDBParams() map[string]string {

	type config struct {
		DBHost string
		DBPort string
		DBName string
		DBUser string
		DBPass string
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

	configMap := make(map[string]string)
	configMap["DBHost"] = configuration.DBHost
	configMap["DBPort"] = configuration.DBPort
	configMap["DBName"] = configuration.DBName
	configMap["DBUser"] = configuration.DBUser
	configMap["DBPass"] = configuration.DBPass

	return configMap

}
