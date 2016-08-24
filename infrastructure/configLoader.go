package infrastructure

import (
	"github.com/naoina/toml"
	"io/ioutil"
)

// Configuration struct holding different options
type Configuration struct {
	Port       int
	DbPassword string
}

// LoadConfigurations loads configurations from pubkeymanager.conf and then returns a pointer
func LoadConfigurations() *Configuration {
	file, err := ioutil.ReadFile("pubkeymanager.conf")
	if err != nil {
		panic(err)
	}
	config := &Configuration{}
	if err := toml.Unmarshal(file, &config); err != nil {
		panic(err)
	}
	return config
}
