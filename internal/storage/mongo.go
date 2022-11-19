package storage

import (
	"context"

	"github.com/acool-kaz/towerOnline/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMognoDb(conf *config.Config) (*mongo.Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.Mongo.Uri))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}
	return client.Database("tower_online"), nil
}
