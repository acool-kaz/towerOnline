package handler

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/acool-kaz/towerOnline/internal/models"
	"github.com/acool-kaz/towerOnline/internal/service"
	"gopkg.in/telebot.v3"
)

func (h *Handler) startHandler(c telebot.Context) error {
	if c.Chat().Type == telebot.ChatGroup {
		return c.Send("Ку, это бот по Башне", h.groupMainMenu)
	} else if c.Chat().Type == telebot.ChatPrivate {
		user := models.User{
			Username:   c.Sender().Username,
			FirstName:  c.Sender().FirstName,
			LastName:   c.Sender().LastName,
			TelegramId: c.Sender().ID,
		}
		if err := h.service.Create(user); err != nil {
			h.logger.Info(err)
			return c.Send("Снова привет, " + c.Sender().Username)
		}
		return c.Send("Ку, " + c.Sender().Username)
	}
	return nil
}

func (h *Handler) ruleHandler(c telebot.Context) error {
	r := &telebot.Poll{
		Question: "text",
		Options: []telebot.PollOption{
			{Text: "1"},
			{Text: "2"},
			{Text: "3"},
		},
		OpenPeriod: 10,
	}

	msg, err := r.Send(h.bot, c.Recipient(), nil)

	time.Sleep(10 * time.Second)
	fmt.Println(msg.Poll.VoterCount)
	return err

}

func (h *Handler) startNewGameHandler(c telebot.Context) error {
	inline := &telebot.ReplyMarkup{}
	inline.Inline(
		telebot.Row{inline.Data("Хочу играть", "add")},
		telebot.Row{inline.Data("Хочу выйти", "leave")},
		telebot.Row{inline.Data("Начинаем", "start")},
		telebot.Row{inline.Data("Отменить", "delete")},
	)
	game := models.Game{
		GroupChatId: c.Chat().ID,
	}
	if err := h.service.Game.CreateGame(game); err != nil {
		h.logger.Info(err)
		return c.Send("Набор в игру уже начался!")
	}
	return c.Send(fmt.Sprintf("Начинаем набор! Начните беседу со мной @%s", h.bot.Me.Username), inline)
}

func (h *Handler) onCallBackHandler(c telebot.Context) error {
	callback := c.Callback()
	if strings.Contains(callback.Data, "poisoner-") {
		h.channel <- strings.Split(callback.Data, "poisoner-")[1]
	} else if strings.Contains(callback.Data, "demon-") {
		h.channel <- strings.Split(callback.Data, "demon-")[1]
	} else if strings.Contains(callback.Data, "fortune-") {
		h.channel <- strings.Split(callback.Data, "fortune-")[1]
	}
	switch callback.Data {
	case "\fadd":
		user, err := h.service.User.GetOne(callback.Sender.ID)
		if err != nil {
			return err
		}
		if err := h.service.Game.JoinGame(user, callback.Message.Chat.ID); err != nil {
			if errors.Is(err, service.ErrGame) {
				h.logger.Info(err)
				if _, err := h.bot.Send(callback.Sender, "Ты уже в игре!"); err != nil {
					return err
				}
				return nil
			}
			if errors.Is(err, service.ErrPlayersTooMany) {
				h.logger.Info(err)
				if _, err := h.bot.Send(callback.Sender, "Слишком много людей. ББ."); err != nil {
					return err
				}
				return nil
			}
			return err
		}
		if _, err := h.bot.Send(callback.Sender, "Ты в игре!"); err != nil {
			return err
		}
	case "\fleave":
		user, err := h.service.User.GetOne(callback.Sender.ID)
		if err != nil {
			return err
		}
		if err := h.service.Game.LeaveGame(user, callback.Message.Chat.ID); err != nil {
			if errors.Is(err, service.ErrGame) {
				h.logger.Info(err)
				if _, err := h.bot.Send(callback.Sender, "Тебя и так нет в игре!"); err != nil {
					return err
				}
				return nil
			}
			return err
		}
		if _, err := h.bot.Send(callback.Sender, "Ты вышел из игры!"); err != nil {
			return err
		}
	case "\fstart":
		game, err := h.service.Game.StartNewGame(callback.Message.Chat.ID)
		if err != nil {
			if errors.Is(err, service.ErrPlayersNotEnough) {
				h.logger.Info(err)
				return c.Send("Недостаточно игроков чтобы начать!")
			}
			return err
		}
		if err := c.Delete(); err != nil {
			return err
		}
		packImg := &telebot.Photo{
			File: telebot.FromDisk("packs/1.png"),
		}
		for _, p := range game.Players {
			if _, err := h.bot.Send(&telebot.User{ID: p.User.TelegramId}, "Игра начинается!"); err != nil {
				return err
			}
			if _, err := h.bot.SendAlbum(&telebot.User{ID: p.User.TelegramId}, telebot.Album{packImg}); err != nil {
				return err
			}
		}
		if err := h.service.Game.SetRoles(&game); err != nil {
			if err := h.service.Game.DeleteGame(callback.Message.Chat.ID); err != nil {
				return err
			}
			h.logger.Info(err)
			return c.Send("Что-то пошло не так. Сори!")
		}
		for _, p := range game.Players {
			if _, err := h.bot.Send(&telebot.User{ID: p.User.TelegramId}, fmt.Sprintf("Твоя роль - %s\nОписание роли - %s", p.Role, p.RoleDescription)); err != nil {
				return err
			}
		}
		if err := c.Send("Игра начинается!"); err != nil {
			return err
		}
		if err := h.service.FirstPack.Start(game, h.bot); err != nil {
			if err := h.service.Game.DeleteGame(game.GroupChatId); err != nil {
				return err
			}
			return err
		}
		if err := c.Send("Игра закончилась!"); err != nil {
			return err
		}
		// if err := h.service.Game.DeleteGame(game.GroupChatId); err != nil {
		// 	return err
		// }
		return nil
	case "\fdelete":
		if err := c.Delete(); err != nil {
			return err
		}
		if err := h.service.Game.DeleteGame(callback.Message.Chat.ID); err != nil {
			return err
		}
		return c.Send("Игра отменена!")
	}
	return nil
}
