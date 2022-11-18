package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	AppConfig struct {
		LogLevel string `env:"LOG-LEVEL" env-default:"trace"`
	}
	Telegram struct {
		Token string `env:"TG-TOKEN" env-required:"true"`
	}
	GameSet []GameConfig `json:"game_set"`
}

type GameConfig struct {
	PlayerCount int `json:"player_count"`
	PlayerSet   struct {
		Townfolks int `json:"townfolks"`
		Outsiders int `json:"outsiders"`
		Minions   int `json:"minions"`
		Demon     int `json:"demon"`
	}
}

var config *Config
var once sync.Once

func GetConfig(path string) *Config {
	once.Do(func() {
		config = &Config{}
		if err := cleanenv.ReadEnv(config); err != nil {
			helpText := "Tower Clock Online Telegram Bot"
			help, _ := cleanenv.GetDescription(config, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
		if err := cleanenv.ReadConfig(path, config); err != nil {
			helpText := "Tower Clock Online Telegram Bot"
			help, _ := cleanenv.GetDescription(config, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return config
}
