package main

import (
	"FluxNodeStats/cli"
	"FluxNodeStats/config"
	"FluxNodeStats/server"
	"FluxNodeStats/utils"
)

func main() {
	cliFilled := cli.Cli()
	config.LoadYamlConfig(cliFilled.FilePathConfig)
	utils.CreateNodes()
	server.GoGinServer()
}
