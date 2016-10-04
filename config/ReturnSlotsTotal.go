package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

func ReturnSlotsTotal() int {

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

	slotsTotalInt, err := strconv.Atoi(c.SlotsTotal)
	if err != nil {
		log.Print(err)
	}

	return slotsTotalInt

}
