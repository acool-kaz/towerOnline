package service

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/acool-kaz/towerOnline/internal/config"
	"github.com/acool-kaz/towerOnline/internal/models"
	"github.com/acool-kaz/towerOnline/internal/storage"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrGame             = errors.New("game error")
	ErrPlayersNotEnough = errors.New("not enough players to start game")
	ErrPlayersTooMany   = errors.New("too many players in game")
)

type Game interface {
	CreateGame(game models.Game) error
	JoinGame(user models.User, groupChatId int64) error
	LeaveGame(user models.User, groupChatId int64) error
	DeleteGame(groupChatId int64) error
	StartNewGame(groupChatId int64) (models.Game, error)
	SetRoles(game *models.Game) error
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
		if p.User.TelegramId == user.TelegramId {
			return fmt.Errorf("%w: user exist in game room", ErrGame)
		}
	}
	game.Players = append(game.Players, models.Player{User: user})
	game.Players = append(game.Players, models.Player{User: user})
	game.Players = append(game.Players, models.Player{User: user})
	game.Players = append(game.Players, models.Player{User: user})
	game.Players = append(game.Players, models.Player{User: user})
	game.Players = append(game.Players, models.Player{User: user})
	return s.stor.ChangePlayers(game)
}

func (s *GameService) LeaveGame(user models.User, groupChatId int64) error {
	game, err := s.stor.GetOne(groupChatId)
	if err != nil {
		return err
	}
	for i, p := range game.Players {
		if p.User.TelegramId == user.TelegramId {
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
	if len(game.Players) < s.config.Game.GameSet[0].PlayerCount {
		return models.Game{}, ErrPlayersNotEnough
	}
	if len(game.Players) > s.config.Game.GameSet[len(s.config.Game.GameSet)-1].PlayerCount {
		return models.Game{}, ErrPlayersTooMany
	}
	return game, nil
}

func (s *GameService) SetRoles(game *models.Game) error {
	for _, set := range s.config.Game.GameSet {
		if set.PlayerCount == len(game.Players) {
			randPlayersIndex := rand.Perm(set.PlayerCount)
			i := 0
			// random demons
			randDemons := rand.Perm(len(s.config.Game.FirstPack.Demons))
			for _, demonsIndex := range randDemons[:set.PlayerSet.Demon] {
				game.Players[randPlayersIndex[i]].Role = s.config.Game.FirstPack.Demons[demonsIndex].Role
				i++
			}

			// random minions
			randMinions := rand.Perm(len(s.config.Game.FirstPack.Minions))
			for _, minionsIndex := range randMinions[:set.PlayerSet.Minions] {
				game.Players[randPlayersIndex[i]].Role = s.config.Game.FirstPack.Minions[minionsIndex].Role
				i++
			}

			if findRole(game.Players, "Baron") {
				set.PlayerSet.Outsiders += 2
				set.PlayerSet.Townfolks -= 2
			}

			// random outsiders
			randOutsiders := rand.Perm(len(s.config.Game.FirstPack.Outsiders))
			for _, outsidersIndex := range randOutsiders[:set.PlayerSet.Outsiders] {
				game.Players[randPlayersIndex[i]].Role = s.config.Game.FirstPack.Outsiders[outsidersIndex].Role
				i++
			}

			// random townfolks
			randTownfolks := rand.Perm(len(s.config.Game.FirstPack.Townsfolks))
			for _, townFolksIndex := range randTownfolks[:set.PlayerSet.Townfolks] {
				game.Players[randPlayersIndex[i]].Role = s.config.Game.FirstPack.Townsfolks[townFolksIndex].Role
				i++
			}
		}
	}
	return nil
}

func findRole(allPlayers []models.Player, role string) bool {
	for _, p := range allPlayers {
		if p.Role == role {
			return true
		}
	}
	return false
}