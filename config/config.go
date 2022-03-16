package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Cfg *Config

type Config struct {
	ApiPort    int    `yaml:"api_port"`
	ApiLogFile string `yaml:"api_log_file"`
	ApiAdress  string `yaml:"api_adress"`
}

func LoadYamlConfig(ConfigFilePath string) {
	t := Config{}
	data, err := ioutil.ReadFile(ConfigFilePath)
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = yaml.Unmarshal(data, &t)
	if err != nil {
		log.Fatalln(err.Error())
	}
	Cfg = &t
}
