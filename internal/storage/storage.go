package storage

import "go.mongodb.org/mongo-driver/mongo"

type Storage struct {
	User
	Game
}

func NewStorage(db *mongo.Database) *Storage {
	return &Storage{
		User: newUserStorage(db),
		Game: newGameStorage(db),
	}
}
