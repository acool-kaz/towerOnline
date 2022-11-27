package models

type Player struct {
	User               User   `bson:"user"`
	Role               string `bson:"role"`
	RoleDescription    string `bson:"role_description"`
	SubRole            string `bson:"sub_role"`
	SubRoleDescription string `bson:"sub_role_description"`
	IsPoison           bool   `bson:"is_poison"`
	IsSafe             bool   `bson:"is_safe"`
	IsDead             bool   `bson:"is_dead"`
	DeadVote           bool   `bson:"dead_vote"`
}
