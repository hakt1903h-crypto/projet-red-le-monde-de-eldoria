package monsters

func InitGoblin() Monster {
	return Monster{
		Name:       "Gobelin d'entrainement",
		MaxHP:      40,
		HP:         40,
		Attack:     5,
		Initiative: 10,
		XPReward:   30,
		GoldReward: 25,
	}
}

func GoblinPattern(turn int) int {
	if turn%3 == 0 {
		return 10
	}

	return 5
}
