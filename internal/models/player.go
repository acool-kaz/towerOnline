package models

type Player struct {
	User               User   `bson:"user" json:"user,omitempty"`
	Role               string `bson:"role" json:"role,omitempty"`
	RoleDescription    string `bson:"role_description" json:"role_description,omitempty"`
	SubRole            string `bson:"sub_role" json:"sub_role,omitempty"`
	SubRoleDescription string `bson:"sub_role_description" json:"sub_role_description,omitempty"`
	FortuneMark        bool   `bson:"fortune_mark" json:"fortune_mark,omitempty"`
	IsPoison           bool   `bson:"is_poison" json:"is_poison,omitempty"`
	IsSafe             bool   `bson:"is_safe" json:"is_safe,omitempty"`
	IsDead             bool   `bson:"is_dead" json:"is_dead,omitempty"`
	DeadVote           bool   `bson:"dead_vote" json:"dead_vote,omitempty"`
}
