package models

type GameState struct {
	Player            Character
	CurrentZone       string
	CurrentScene      string
	MemoryFragments   []string
	DefeatedBosses    []string
	StoryProgress     string
	IsGameFinished    bool
	TrainingCompleted bool
	VisitedZones      []string
}
