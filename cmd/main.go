package main

import (
	"evenApi/api"
	"evenApi/config"
	"evenApi/logger"
	"evenApi/service"
	"fmt"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")
	fmt.Println("after logger")

	serviceManager, err := service.NewserviceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	apiToRun := api.New(api.Options{
		Cfg:               cfg,
		Log:               log,
		ServiceToDoClient: serviceManager,
	})

	if err := apiToRun.Run(cfg.HTTPPort); err != nil {
		log.Error("Cant running gin engine", logger.Error(err))
		return
	}

}
