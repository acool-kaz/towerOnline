package storage

import (
	"context"

	"github.com/acool-kaz/towerOnline/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User interface {
	Create(user models.User) error
	GetOne(telegramId int64) (models.User, error)
}

type UserStorage struct {
	db *mongo.Database
}

func newUserStorage(db *mongo.Database) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (s *UserStorage) Create(user models.User) error {
	coll := s.db.Collection("users")
	if _, err := coll.InsertOne(context.TODO(), &user); err != nil {
		return err
	}
	return nil
}

func (s *UserStorage) GetOne(telegramId int64) (models.User, error) {
	var user models.User
	coll := s.db.Collection("users")
	if err := coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "telegram_id", Value: telegramId}}).Decode(&user); err != nil {
		return models.User{}, err
	}
	return user, nil
}
