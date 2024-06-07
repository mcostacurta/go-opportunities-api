package main

import (
	"github.com/mcostacurta/go-opportunities-api/config"
	"github.com/mcostacurta/go-opportunities-api/router"
)

var (
	logger config.Logger
)

func main() {

	logger = *config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	router.Initialize() // listen and serve on 0.0.0.0:8080
}
