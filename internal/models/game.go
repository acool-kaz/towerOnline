package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Game struct {
	ID          primitive.ObjectID `bson:"_id"`
	Players     []Player           `bson:"players"`
	GroupChatId int64              `bson:"group_chat_id"`
}
