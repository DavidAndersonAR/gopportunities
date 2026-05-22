package main

import (
	

	"github.com/DavidAndersonAR/gopportunities.git/config"
	"github.com/DavidAndersonAR/gopportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}


	// Initialize router
	router.Initialize()

}