package models

type Player struct {
	User     User   `bson:"user"`
	Role     string `bson:"role"`
	SubRole  string `bson:"sub_role"`
	IsPoison bool   `bson:"is_poison"`
	IsSafe   bool   `bson:"is_safe"`
	IsDead   bool   `bson:"is_dead"`
	DeadVote bool   `bson:"dead_vote"`
}
