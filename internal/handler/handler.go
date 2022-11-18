package handler

import (
	"github.com/acool-kaz/towerOnline/internal/config"
	"github.com/acool-kaz/towerOnline/internal/service"
	"github.com/acool-kaz/towerOnline/pkg/logger"
	"gopkg.in/telebot.v3"
)

var (
	groupMainMenu = &telebot.ReplyMarkup{}
	ruleBtn       = groupMainMenu.Data("Правила", "rules")
	startGameBtn  = groupMainMenu.Data("Начать новую игру", "start_new_game")
)

type Handler struct {
	bot           *telebot.Bot
	config        *config.Config
	logger        *logger.Logger
	service       *service.Service
	groupMainMenu *telebot.ReplyMarkup
}

func NewHandler(b *telebot.Bot, c *config.Config, l *logger.Logger, srv *service.Service) *Handler {
	return &Handler{
		bot:           b,
		config:        c,
		logger:        l,
		service:       srv,
		groupMainMenu: groupMainMenu,
	}
}

func (h *Handler) StartBot() {
	groupMainMenu.Inline(telebot.Row{ruleBtn}, telebot.Row{startGameBtn})

	h.bot.Handle("/start", h.startHandler)
	h.bot.Handle(&ruleBtn, h.ruleHandler)
	h.bot.Handle(&startGameBtn, h.startNewGameHandler)
	h.bot.Handle(telebot.OnCallback, h.onCallBackHandler)

	h.bot.Start()
}
