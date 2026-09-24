package models

type DialogueChoice struct {
	Text             string
	NextDialogueID   string
	SetStoryProgress string
	MemoryID         string
}

type Dialogue struct {
	ID      string
	Text    string
	Choices []DialogueChoice
}

type NPC struct {
	Name      string
	Dialogues []Dialogue
	State     int
}
