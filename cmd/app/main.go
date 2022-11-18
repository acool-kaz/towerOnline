package main

import (
	"log"

	"github.com/acool-kaz/towerOnline/internal/app"
	"github.com/acool-kaz/towerOnline/internal/config"
	"github.com/acool-kaz/towerOnline/pkg/logger"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
}

func main() {
	log.Print("confing init")
	cfg := config.GetConfig("./gameconfig.json")

	log.Print("logger init")
	logger := logger.GetLogger(cfg.AppConfig.LogLevel)

	a := app.NewApp(cfg, logger)
	a.Run()
}
