package service

import (
	"fmt"
	"math/rand"

	"github.com/acool-kaz/towerOnline/internal/config"
	"github.com/acool-kaz/towerOnline/internal/models"
	"github.com/acool-kaz/towerOnline/internal/storage"
	"gopkg.in/telebot.v3"
)

type FirstPack interface {
	Start(game models.Game, bot *telebot.Bot) error
	zeroNight(game models.Game, bot *telebot.Bot, firstPack config.Pack) error
	otherNight(game models.Game, bot *telebot.Bot) error
	poisoner(game models.Game, bot *telebot.Bot) error
	imp(game models.Game, bot *telebot.Bot) error
	washerwoman(game models.Game, bot *telebot.Bot) error
	librarian(game models.Game, bot *telebot.Bot) error
	investigator(game models.Game, bot *telebot.Bot) error
	empath(game models.Game, bot *telebot.Bot) error
	chef(game models.Game, bot *telebot.Bot) error
}

type FirstPackService struct {
	stor    storage.Game
	cfg     *config.Config
	channel chan string
}

func newFirstPackService(stor storage.Game, cfg *config.Config, channel chan string) *FirstPackService {
	return &FirstPackService{
		stor:    stor,
		cfg:     cfg,
		channel: channel,
	}
}

func (s *FirstPackService) Start(game models.Game, bot *telebot.Bot) error {
	var err error
	if err = s.zeroNight(game, bot, s.cfg.Game.FirstPack); err != nil {
		return err
	}
	// for {
	// 	if err := s.otherNight(game, bot); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func (s *FirstPackService) zeroNight(game models.Game, bot *telebot.Bot, firstPack config.Pack) error {
	var err error
	if err = minionInfo(game, bot); err != nil {
		return err
	}
	if err = demonInfo(game, bot, firstPack); err != nil {
		return err
	}
	if err = s.poisoner(game, bot); err != nil {
		return err
	}
	if err = s.washerwoman(game, bot); err != nil {
		return err
	}
	if err = s.librarian(game, bot); err != nil {
		return err
	}
	if err = s.investigator(game, bot); err != nil {
		return err
	}
	if err = s.chef(game, bot); err != nil {
		return err
	}
	if err = s.empath(game, bot); err != nil {
		return err
	}
	return nil
}

func (s *FirstPackService) otherNight(game models.Game, bot *telebot.Bot) error {
	return nil
}

func minionInfo(game models.Game, bot *telebot.Bot) error {
	for _, p := range game.Minions {
		if _, err := bot.Send(&telebot.User{ID: p.User.TelegramId}, fmt.Sprintf("Твой демон - %s", game.Demons[0].User.Username)); err != nil {
			return err
		}
	}
	return nil
}

func demonInfo(game models.Game, bot *telebot.Bot, firstPack config.Pack) error {
	for _, p := range game.Minions {
		if _, err := bot.Send(&telebot.User{ID: game.Demons[0].User.TelegramId}, fmt.Sprintf("Твои миньоны - %s", p.User.Username)); err != nil {
			return err
		}
	}
	var allRoles []string
	for _, r := range firstPack.Townsfolks {
		allRoles = append(allRoles, r.Role)
	}
	for _, r := range firstPack.Outsiders {
		allRoles = append(allRoles, r.Role)
	}
	var notInGame []string
	for _, role := range allRoles {
		if !isInArray(role, game.Players) {
			notInGame = append(notInGame, role)
		}
	}
	randNotInGame := rand.Perm(len(notInGame))
	if _, err := bot.Send(&telebot.User{ID: game.Demons[0].User.TelegramId}, "Свободные роли:"); err != nil {
		return err
	}
	for _, index := range randNotInGame[:3] {
		if _, err := bot.Send(&telebot.User{ID: game.Demons[0].User.TelegramId}, notInGame[index]); err != nil {
			return err
		}
	}
	return nil
}

func isInArray(str string, arr []models.Player) bool {
	for _, s := range arr {
		if str == s.Role {
			return true
		}
	}
	return false
}

func (s *FirstPackService) poisoner(game models.Game, bot *telebot.Bot) error {
	inline := &telebot.ReplyMarkup{}
	rows := []telebot.Row{}
	for _, p := range game.Players {
		rows = append(rows, telebot.Row{inline.Data(p.User.Username, "poisoner-"+p.User.Username)})
	}
	inline.Inline(rows...)
	for _, p := range game.Minions {
		if p.Role == "Poisoner" {
			if _, err := bot.Send(&telebot.User{ID: p.User.TelegramId}, "Можешь отравить одного!", inline); err != nil {
				return err
			}
			poisonUser := <-s.channel
			for i, p := range game.Players {
				if p.User.Username == poisonUser {
					game.Players[i].IsPoison = true
				}
			}
			if err := s.stor.ChangePlayers(game); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *FirstPackService) washerwoman(game models.Game, bot *telebot.Bot) error {
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

func (s *FirstPackService) librarian(game models.Game, bot *telebot.Bot) error {
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

func (s *FirstPackService) investigator(game models.Game, bot *telebot.Bot) error {
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

func (s *FirstPackService) chef(game models.Game, bot *telebot.Bot) error {
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

func (s *FirstPackService) empath(game models.Game, bot *telebot.Bot) error {
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

func (s *FirstPackService) imp(game models.Game, bot *telebot.Bot) error {
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
