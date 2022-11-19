package models

type Game struct {
	GroupChatId int64    `bson:"group_chat_id"`
	Players     []Player `bson:"players"`
}
