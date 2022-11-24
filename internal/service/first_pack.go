package service

import (
	"github.com/acool-kaz/towerOnline/internal/config"
	"github.com/acool-kaz/towerOnline/internal/models"
	"gopkg.in/telebot.v3"
)

type FirstPack interface {
	Start(game models.Game, bot *telebot.Bot) error
	Washerwoman(game models.Game, bot *telebot.Bot) error
}

type FirstPackService struct {
	cfg *config.Config
}

func newFirstPackService(cfg *config.Config) *FirstPackService {
	return &FirstPackService{
		cfg: cfg,
	}
}

func (s *FirstPackService) Start(game models.Game, bot *telebot.Bot) error {
	return nil
}

func (s *FirstPackService) Washerwoman(game models.Game, bot *telebot.Bot) error {

	return nil
}
