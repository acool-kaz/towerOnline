package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	Username   string             `bson:"username"`
	FirstName  string             `bson:"first_name"`
	LastName   string             `bson:"last_name"`
	TelegramId int64              `bson:"telegram_id"`
}
