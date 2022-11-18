package service

import (
	"errors"
	"fmt"

	"github.com/acool-kaz/towerOnline/internal/config"
	"github.com/acool-kaz/towerOnline/internal/models"
	"github.com/acool-kaz/towerOnline/internal/storage"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrGame = errors.New("game error")
var ErrPlayersNotEnough = errors.New("not enough players to start game")
var ErrPlayersTooMany = errors.New("too many players in game")

type Game interface {
	CreateGame(game models.Game) error
	JoinGame(user models.User, groupChatId int64) error
	LeaveGame(user models.User, groupChatId int64) error
	DeleteGame(groupChatId int64) error
	StartNewGame(groupChatId int64) (models.Game, error)
}

type GameService struct {
	stor   storage.Game
	config *config.Config
}

func newGameService(stor storage.Game, c *config.Config) *GameService {
	return &GameService{
		stor:   stor,
		config: c,
	}
}

func (s *GameService) CreateGame(game models.Game) error {
	if _, err := s.stor.GetOne(game.GroupChatId); err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return err
		}
	}
	return s.stor.CreateGame(game)
}

func (s *GameService) JoinGame(user models.User, groupChatId int64) error {
	game, err := s.stor.GetOne(groupChatId)
	if err != nil {
		return err
	}
	for _, p := range game.Players {
		if p.User.ID == user.ID {
			return fmt.Errorf("%w: user exist in game room", ErrGame)
		}
	}
	game.Players = append(game.Players, models.Player{User: user})
	return s.stor.ChangePlayers(game)
}

func (s *GameService) LeaveGame(user models.User, groupChatId int64) error {
	game, err := s.stor.GetOne(groupChatId)
	if err != nil {
		return err
	}
	for i, p := range game.Players {
		if p.User.ID == user.ID {
			game.Players = append(game.Players[:i], game.Players[i+1:]...)
			break
		}
	}
	return s.stor.ChangePlayers(game)
}

func (s *GameService) DeleteGame(groupChatId int64) error {
	return s.stor.DeleteGame(groupChatId)
}

func (s *GameService) StartNewGame(groupChatId int64) (models.Game, error) {
	game, err := s.stor.GetOne(groupChatId)
	if err != nil {
		return models.Game{}, err
	}
	// if len(game.Players) < s.config.GameSet[0].PlayerCount {
	// 	return models.Game{}, ErrPlayersNotEnough
	// }
	// if len(game.Players) > s.config.GameSet[len(s.config.GameSet)-1].PlayerCount {
	// 	return models.Game{}, ErrPlayersTooMany
	// }
	return game, nil
}
