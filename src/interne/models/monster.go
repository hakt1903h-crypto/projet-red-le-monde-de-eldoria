package models

type Monster struct {
	Name             string
	MaxHP            int
	CurrentHP        int
	Attack           int
	Initiative       int
	ExperienceReward int
	Defense          int
	IsBoss           bool
	BossType         string
}
