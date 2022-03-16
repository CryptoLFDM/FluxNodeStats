package cli

import (
	"flag"
)

type ConfigStruct struct {
	FilePathConfig string
}

func Cli() ConfigStruct {
	config := ConfigStruct{}
	flag.StringVar(&config.FilePathConfig, "config", "config/config.yml", "file config path")
	flag.Parse() // This will parse all the arguments from the terminal
	return config
}
