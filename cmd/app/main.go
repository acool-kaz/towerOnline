package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/acool-kaz/towerOnline/internal/app"
	"github.com/acool-kaz/towerOnline/internal/config"
	"github.com/acool-kaz/towerOnline/pkg/logger"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Print("confing init")
	cfg := config.GetConfig()

	log.Print("logger init")
	logger := logger.GetLogger(cfg.AppConfig.LogLevel)

	a := app.NewApp(cfg, logger)
	a.Run()
}
