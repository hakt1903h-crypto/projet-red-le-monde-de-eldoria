package player

import "fmt"

// Gains de statistiques à chaque montée de niveau.
const (
	XPPerLevelGrowth = 50 // +50 XP requis par niveau supplémentaire
	HPPerLevel       = 10
	ManaPerLevel     = 5
	AttackPerLevel   = 2
)

// GainXP ajoute de l'expérience au personnage et gère les montées de niveau.
// L'excédent d'XP est conservé (ex : besoin 100, XP 90, gain +30 -> niveau +1, reste 20).
func GainXP(character *Character, amount int) {
	if amount <= 0 {
		return
	}

	character.XP += amount

	for character.XP >= character.MaxXP {
		character.XP -= character.MaxXP
		LevelUp(character)
	}
}

// LevelUp fait monter le personnage d'un niveau et augmente ses statistiques.
func LevelUp(character *Character) {
	character.Level++
	character.MaxXP += XPPerLevelGrowth

	character.MaxHP += HPPerLevel
	character.HP += HPPerLevel

	character.MaxMana += ManaPerLevel
	character.Mana += ManaPerLevel

	character.Attack += AttackPerLevel

	fmt.Printf("Niveau supérieur ! Vous êtes maintenant niveau %d.\n", character.Level)
	fmt.Printf("PV : %d/%d | Mana : %d/%d | Attaque : %d\n",
		character.HP, character.MaxHP, character.Mana, character.MaxMana, character.Attack)
}

// IsAlive indique si le personnage est encore en vie.
func IsAlive(character Character) bool {
	return character.HP > 0
}

// Revive ranime un personnage mort avec 50 % de ses PV maximum.
// Retourne false si le personnage n'était pas mort.
func Revive(character *Character) bool {
	if character.HP > 0 {
		return false
	}

	character.HP = character.MaxHP / 2
	fmt.Printf("Vous revenez à la vie avec %d/%d PV.\n", character.HP, character.MaxHP)
	return true
}

// Heal soigne le personnage sans dépasser ses PV maximum.
func Heal(character *Character, amount int) {
	if amount <= 0 {
		return
	}

	character.HP += amount
	if character.HP > character.MaxHP {
		character.HP = character.MaxHP
	}
}
