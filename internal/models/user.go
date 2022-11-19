package models

type User struct {
	Username   string `bson:"username"`
	FirstName  string `bson:"first_name"`
	LastName   string `bson:"last_name"`
	TelegramId int64  `bson:"telegram_id"`
}
