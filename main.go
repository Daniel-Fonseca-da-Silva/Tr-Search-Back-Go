package main

import (
	"github.com/Daniel-Fonseca-da-Silva/Tr-Search-Back-Go/config"
	"github.com/Daniel-Fonseca-da-Silva/Tr-Search-Back-Go/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()

	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
