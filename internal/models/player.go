package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	ID       primitive.ObjectID `bson:"_id"`
	User     User               `bson:"user"`
	Role     string             `bson:"role"`
	IsDead   bool               `bson:"is_dead"`
	DeadVote bool               `bson:"dead_vote"`
}
