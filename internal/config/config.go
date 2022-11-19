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
	Mongo struct {
		Uri string `env:"URI" env-required:"true"`
	}
	Game struct {
		GameSet   []GameConfig `json:"game_set"`
		FirstPack Pack         `json:"first_pack"`
	} `json:"game"`
}

type Pack struct {
	Townsfolks []Profile `json:"townsfolks"`
	Outsiders  []Profile `json:"outsiders"`
	Minions    []Profile `json:"minions"`
	Demons     []Profile `json:"demons"`
}

type Profile struct {
	Role        string `json:"role"`
	Description string `json:"description"`
}

type GameConfig struct {
	PlayerCount int `json:"player_count"`
	PlayerSet   struct {
		Townfolks int `json:"townfolks"`
		Outsiders int `json:"outsiders"`
		Minions   int `json:"minions"`
		Demon     int `json:"demon"`
	} `json:"player_set"`
}

var (
	config *Config
	once   sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{}
		if err := cleanenv.ReadEnv(config); err != nil {
			helpText := "Tower Clock Online Telegram Bot"
			help, _ := cleanenv.GetDescription(config, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
		if err := cleanenv.ReadConfig("./gameconfig.json", config); err != nil {
			helpText := "Tower Clock Online Telegram Bot"
			help, _ := cleanenv.GetDescription(config, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
		if err := cleanenv.ReadConfig("./1pack.json", config); err != nil {
			helpText := "Tower Clock Online Telegram Bot"
			help, _ := cleanenv.GetDescription(config, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return config
}
