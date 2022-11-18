package service

import (
	"github.com/acool-kaz/towerOnline/internal/config"
	"github.com/acool-kaz/towerOnline/internal/storage"
)

type Service struct {
	User
	Game
}

func NewService(storage *storage.Storage, c *config.Config) *Service {
	return &Service{
		User: newUserService(storage.User),
		Game: newGameService(storage.Game, c),
	}
}
