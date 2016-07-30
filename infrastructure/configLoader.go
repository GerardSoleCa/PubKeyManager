package infrastructure

import (
	"github.com/naoina/toml"
	"io/ioutil"
)

type Configuration struct {
	Port int
	DbPassword string
}

func LoadConfigurations() (*Configuration) {
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
