package models

type Game struct {
	GroupChatId int64    `bson:"group_chat_id"`
	Townfolks   []Player `bson:"townfolks"`
	Outsiders   []Player `bson:"outsiders"`
	Minions     []Player `bson:"minions"`
	Demons      []Player `bson:"demons"`
	Players     []Player `bson:"players"`
}
