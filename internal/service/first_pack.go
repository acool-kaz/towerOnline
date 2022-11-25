package service

import (
	"fmt"
	"math/rand"

	"github.com/acool-kaz/towerOnline/internal/config"
	"github.com/acool-kaz/towerOnline/internal/models"
	"gopkg.in/telebot.v3"
)

type FirstPack interface {
	Start(game models.Game, bot *telebot.Bot) error
	Washerwoman(game models.Game, bot *telebot.Bot) error
	Librarian(game models.Game, bot *telebot.Bot) error
	Investigator(game models.Game, bot *telebot.Bot) error
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
	s.Washerwoman(game, bot)
	s.Librarian(game, bot)
	s.Investigator(game, bot)
	return nil
}

func (s *FirstPackService) Washerwoman(game models.Game, bot *telebot.Bot) error {
	var (
		info            string
		truePlayerInfo  models.Player
		falsePlayerInfo models.Player
	)
	for {
		truePlayerInfo = game.Townfolks[rand.Intn(len(game.Townfolks))]
		if truePlayerInfo.Role != "Washerwoman" {
			break
		}
	}
	for {
		falsePlayerInfo = game.Players[rand.Intn(len(game.Players))]
		if truePlayerInfo.Role != falsePlayerInfo.Role {
			break
		}
	}
	if rand.Intn(100) < 50 {
		info = fmt.Sprintf("Кто-то из %s, %s является: %s", falsePlayerInfo.User.Username, truePlayerInfo.User.Username, truePlayerInfo.Role)
	} else {
		info = fmt.Sprintf("Кто-то из %s, %s является: %s", truePlayerInfo.User.Username, falsePlayerInfo.User.Username, truePlayerInfo.Role)
	}
	for _, p := range game.Townfolks {
		if p.Role == "Washerwoman" {
			if _, err := bot.Send(&telebot.User{ID: p.User.TelegramId}, info); err != nil {
				return err
			}
			break
		}
	}
	fmt.Println(info)
	return nil
}

func (s *FirstPackService) Librarian(game models.Game, bot *telebot.Bot) error {
	var (
		info            string
		truePlayerInfo  models.Player
		falsePlayerInfo models.Player
	)
	truePlayerInfo = game.Outsiders[rand.Intn(len(game.Outsiders))]
	for {
		falsePlayerInfo = game.Players[rand.Intn(len(game.Players))]
		if truePlayerInfo.Role != falsePlayerInfo.Role {
			break
		}
	}
	if rand.Intn(100) < 50 {
		info = fmt.Sprintf("Кто-то из %s, %s является: %s", falsePlayerInfo.User.Username, truePlayerInfo.User.Username, truePlayerInfo.Role)
	} else {
		info = fmt.Sprintf("Кто-то из %s, %s является: %s", truePlayerInfo.User.Username, falsePlayerInfo.User.Username, truePlayerInfo.Role)
	}
	for _, p := range game.Townfolks {
		if p.Role == "Librarian" {
			if _, err := bot.Send(&telebot.User{ID: p.User.TelegramId}, info); err != nil {
				return err
			}
			break
		}
	}
	fmt.Println(info)
	return nil
}

func (s *FirstPackService) Investigator(game models.Game, bot *telebot.Bot) error {
	var (
		info            string
		truePlayerInfo  models.Player
		falsePlayerInfo models.Player
	)
	truePlayerInfo = game.Minions[rand.Intn(len(game.Minions))]
	for {
		falsePlayerInfo = game.Players[rand.Intn(len(game.Players))]
		if truePlayerInfo.Role != falsePlayerInfo.Role {
			break
		}
	}
	if rand.Intn(100) < 50 {
		info = fmt.Sprintf("Кто-то из %s, %s является: %s", falsePlayerInfo.User.Username, truePlayerInfo.User.Username, truePlayerInfo.Role)
	} else {
		info = fmt.Sprintf("Кто-то из %s, %s является: %s", truePlayerInfo.User.Username, falsePlayerInfo.User.Username, truePlayerInfo.Role)
	}
	for _, p := range game.Townfolks {
		if p.Role == "Investigator" {
			if _, err := bot.Send(&telebot.User{ID: p.User.TelegramId}, info); err != nil {
				return err
			}
			break
		}
	}
	fmt.Println(info)
	return nil
}
