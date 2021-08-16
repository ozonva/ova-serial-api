package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	test int
}

func (c Config) String() string {
	return fmt.Sprintf("{test: %d}", c.test)
}

func UpdateConfig(path string, config *Config) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func(path string, f *os.File) {
		err := f.Close()
		fmt.Printf("Closing file: %s\n", path)
		if err != nil {
			panic(err)
		}
	}(path, f)

	decoder := json.NewDecoder(f)

	err = decoder.Decode(config)
	if err != nil {
		return err
	}

	return nil
}
