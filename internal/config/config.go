package config

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"path/filepath"
	"sync"
)

type Config struct {
	FilePath string
	data     interface{}
	mutex    sync.Mutex
}

func (config *Config) Update() (ret error) {
	config.mutex.Lock()

	defer func() {
		config.mutex.Unlock()
	}()

	filename, _ := filepath.Abs(config.FilePath)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &config.data)
	if err != nil {
		return err
	}

	return nil
}

func (config *Config) GetData() interface{} {
	return config.data
}
