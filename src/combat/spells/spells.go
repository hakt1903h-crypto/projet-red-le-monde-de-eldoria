package spells

type Spell struct {
	Name   string
	Damage int
	Mana   int
}

var Punch = Spell{
	Name:   "Coup de Poing",
	Damage: 8,
	Mana:   5,
}

var Fireball = Spell{
	Name:   "Boule de Feu",
	Damage: 18,
	Mana:   15,
}
