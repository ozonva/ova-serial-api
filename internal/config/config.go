package config

import (
	"github.com/ghodss/yaml"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type Config struct {
	Data  interface{}
	mutex sync.Mutex
}

func UpdateConfig(path string, config *Config) (ret error) {
	config.mutex.Lock()

	f, err := os.Open(path)
	if err != nil {
		ret = err
	}

	defer func(path string, f *os.File) {
		err := f.Close()
		if err != nil {
			log.Error().Msgf("An error occurred when closing file: %+v", err)
		}
	}(path, f)

	defer func() {
		config.mutex.Unlock()
	}()

	filename, _ := filepath.Abs(path)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		ret = err
	}

	err = yaml.Unmarshal(yamlFile, &config.Data)
	if err != nil {
		ret = err
	}

	return
}
