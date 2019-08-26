package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	CleanJob []CleanJob `yaml:"CleanJob"`
}

type CleanJob struct {
	Cron        string   `yaml:"Cron"`
	Endpoints   []string `yaml:"Endpoints"`
	NamePattern string   `yaml:"NamePattern"`
	DatePattern string   `yaml:"DatePattern"`
	Retention   int      `yaml"Retention"`
}

func ReadConfig(fn string) *Config {
	rv := &Config{}

	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(b, rv)
	if err != nil {
		panic(err)
	}

	return rv
}
