package models

type Player struct {
	User     User   `bson:"user"`
	Role     string `bson:"role"`
	IsDead   bool   `bson:"is_dead"`
	DeadVote bool   `bson:"dead_vote"`
}
