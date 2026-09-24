package models

type Character struct {
	Name              string
	Class             string
	Level             int
	MaxHP             int
	CurrentHP         int
	Inventory         []Item
	InventoryCapacity int
	InventoryUpgrades int
	Money             int
	Spells            []Spell
	Mana              int
	MaxMana           int
	XP                int
	MaxXP             int
	Initiative        int

	LevelHPBonus   int
	LevelManaBonus int
	SilencedTurns  int

	Equipment       Equipment
	MemoryFragments []string
}
