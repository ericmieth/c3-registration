package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

func ReturnSlotsTotal() int {

	type config struct {
		SlotsTotal string
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

	slotsTotalInt, err := strconv.Atoi(configuration.SlotsTotal)
	if err != nil {
		log.Print(err)
	}

	return slotsTotalInt

}
