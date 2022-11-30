package service

import (
	"github.com/acool-kaz/towerOnline/internal/config"
	"github.com/acool-kaz/towerOnline/internal/storage"
)

type Service struct {
	User
	Game
	FirstPack
}

func NewService(storage *storage.Storage, c *config.Config, channel chan string) *Service {
	return &Service{
		User:      newUserService(storage.User),
		Game:      newGameService(storage.Game, c),
		FirstPack: newFirstPackService(storage.Game, c, channel),
	}
}
