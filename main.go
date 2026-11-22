package main

import (
	"github.com/galvao0/fitplan/config"
	"github.com/galvao0/fitplan/router"
)

func main() {
	logger := config.GetLogger("main")

	if err := config.Init(); err != nil {
		logger.Errorf("config initialize error: %v", err.Error())
		return
	}

	router.Init()
}
