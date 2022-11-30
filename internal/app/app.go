package app

import (
	"time"

	"github.com/acool-kaz/towerOnline/internal/config"
	"github.com/acool-kaz/towerOnline/internal/handler"
	"github.com/acool-kaz/towerOnline/internal/service"
	"github.com/acool-kaz/towerOnline/internal/storage"
	"github.com/acool-kaz/towerOnline/pkg/logger"
	"gopkg.in/telebot.v3"
)

type App struct {
	config *config.Config
	logger *logger.Logger
}

func NewApp(c *config.Config, l *logger.Logger) *App {
	return &App{
		config: c,
		logger: l,
	}
}

func (a *App) Run() {
	a.logger.Print("bot init")
	bot, err := a.initBot()
	if err != nil {
		a.logger.Fatal(err)
	}

	a.logger.Print("db init")
	db, err := storage.NewMognoDb(a.config)
	if err != nil {
		a.logger.Fatal(err)
	}
	channel := make(chan string, 1)
	storage := storage.NewStorage(db)
	service := service.NewService(storage, a.config, channel)
	handler := handler.NewHandler(bot, a.config, a.logger, service, channel)

	a.logger.Print("start bot")
	handler.StartBot()
}

func (a *App) initBot() (*telebot.Bot, error) {
	pref := telebot.Settings{
		Token:  a.config.Telegram.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := telebot.NewBot(pref)
	if err != nil {
		return nil, err
	}
	return b, nil
}
