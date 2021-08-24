package main

import (
	"fmt"
	"ova-serial-api/internal/config"
	"time"
)

const configPath = "config/test_config.json"
const intervalSec = 10

func main() {
	fmt.Println("ova-serial-api")

	var cfg config.Config
	for {
		err := config.UpdateConfig(configPath, &cfg)
		if err != nil {
			fmt.Printf("Error occurred: %s\n", nil)
		} else {
			fmt.Printf("Config '%s' updated: %+v\n", configPath, cfg)
		}

		time.Sleep(intervalSec * time.Second)
	}
}
