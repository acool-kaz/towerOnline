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
	Chef(game models.Game, bot *telebot.Bot) error
	Empath(game models.Game, bot *telebot.Bot) error
	Imp(game models.Game, bot *telebot.Bot) error
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
	s.Chef(game, bot)
	s.Empath(game, bot)
	s.Imp(game, bot)
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

func (s *FirstPackService) Chef(game models.Game, bot *telebot.Bot) error {
	var (
		count int
		info  string
	)
	for i := 0; i < len(game.Players); i++ {
		cur := game.Players[i]
		next := game.Players[(i+1)%len(game.Players)]
		if isEvil(cur) && isEvil(next) {
			count++
		}
	}
	if count == 1 {
		info = fmt.Sprintf("Есть %d пара злодеев.", count)
	} else {
		info = fmt.Sprintf("Есть %d пары злодеев.", count)
	}
	for _, p := range game.Townfolks {
		if p.Role == "Chef" {
			if _, err := bot.Send(&telebot.User{ID: p.User.TelegramId}, info); err != nil {
				return err
			}
			break
		}
	}
	fmt.Println(info)
	return nil
}

func (s *FirstPackService) Empath(game models.Game, bot *telebot.Bot) error {
	var (
		count int
		info  string
		prev  models.Player
		next  models.Player
	)
	for i, p := range game.Players {
		if p.Role == "Empath" {
			index := i
			for {
				index--
				if index == -1 {
					index = len(game.Players)
				}
				prev = game.Players[index]
				if !prev.IsDead {
					break
				}
			}
			index = i
			for {
				index++
				index %= len(game.Players)
				next = game.Players[index]
				if !next.IsDead {
					break
				}
			}

			if isEvil(next) {
				count++
			}
			if isEvil(prev) {
				count++
			}
			info = fmt.Sprintf("Кол-во злых соседей: %d", count)

			if _, err := bot.Send(&telebot.User{ID: p.User.TelegramId}, info); err != nil {
				return err
			}
			break
		}
	}
	fmt.Println(info)
	return nil
}

func (s *FirstPackService) Imp(game models.Game, bot *telebot.Bot) error {
	inline := &telebot.ReplyMarkup{}
	rows := []telebot.Row{}
	for _, p := range game.Players {
		rows = append(rows, telebot.Row{inline.Data(p.User.Username, p.User.Username)})
	}
	inline.Inline(rows...)
	for _, p := range game.Demons {
		if _, err := bot.Send(&telebot.User{ID: p.User.TelegramId}, "Можешь убить одного!", inline); err != nil {
			return err
		}
	}
	return nil
}

func isEvil(player models.Player) bool {
	for _, role := range []string{"Poisoner", "Recluse", "Scarlet Woman", "Baron", "Imp"} {
		if player.Role == role {
			return true
		}
	}
	return false
}
