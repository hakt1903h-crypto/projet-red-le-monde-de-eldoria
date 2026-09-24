package models

type Zone struct {
	Name        string
	Description string
	Enemies     []Monster
	NPCs        []NPC
	Connections map[string]string
}
