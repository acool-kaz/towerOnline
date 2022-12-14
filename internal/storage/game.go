package storage

import (
	"context"

	"github.com/acool-kaz/towerOnline/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Game interface {
	CreateGame(game models.Game) error
	GetOne(groupChatId int64) (models.Game, error)
	ChangePlayers(game models.Game) error
	DeleteGame(groupChatId int64) error
}

type GameStorage struct {
	db *mongo.Database
}

func newGameStorage(db *mongo.Database) *GameStorage {
	return &GameStorage{
		db: db,
	}
}

func (s *GameStorage) CreateGame(game models.Game) error {
	coll := s.db.Collection("games")
	if _, err := coll.InsertOne(context.TODO(), &game); err != nil {
		return err
	}
	return nil
}

func (s *GameStorage) GetOne(groupChatId int64) (models.Game, error) {
	var game models.Game
	coll := s.db.Collection("games")
	err := coll.FindOne(context.TODO(), bson.D{
		primitive.E{
			Key:   "group_chat_id",
			Value: groupChatId,
		},
	}).Decode(&game)
	if err != nil {
		return models.Game{}, err
	}
	return game, nil
}

func (s *GameStorage) ChangePlayers(game models.Game) error {
	coll := s.db.Collection("games")
	filter := bson.D{primitive.E{Key: "group_chat_id", Value: game.GroupChatId}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "townfolks", Value: game.Townfolks},
		primitive.E{Key: "outsiders", Value: game.Outsiders},
		primitive.E{Key: "minions", Value: game.Minions},
		primitive.E{Key: "demons", Value: game.Demons},
		primitive.E{Key: "players", Value: game.Players},
	}}}
	return coll.FindOneAndUpdate(context.TODO(), filter, update).Err()
}

func (s *GameStorage) DeleteGame(groupChatId int64) error {
	coll := s.db.Collection("games")
	filter := bson.D{primitive.E{Key: "group_chat_id", Value: groupChatId}}
	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
