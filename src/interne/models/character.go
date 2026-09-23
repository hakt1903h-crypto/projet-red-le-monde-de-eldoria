package models

type Character struct {
	Name            string
	Class           string
	Level           int
	MaxHP           int
	CurrentHP       int
	Inventory       []Item
	Money           int
	Spells          []Spell
	Mana            int
	MaxMana         int
	XP              int
	MaxXP           int
	Initiative      int
	Equipment       Equipment
	MemoryFragments []string
}
