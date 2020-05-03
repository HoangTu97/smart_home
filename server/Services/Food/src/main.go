package main

import (
	"Food/util/logger"
	"Food/config"
)

func main() {
	config.LoadConfigFile()
	config.LoadAppProperties()

	logger.Setup()

	config.SetupDB()
	config.InitResource()
	config.SetupRoutes()
	config.CloseDB()
}
