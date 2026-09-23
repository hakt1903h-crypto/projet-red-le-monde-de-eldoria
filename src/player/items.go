package player

import (
	"eldoria/combat/spells"
	"eldoria/inventory"
	"fmt"
)

const HealthPotionHeal = 50

// UseHealthPotion consomme une potion de vie et soigne le personnage (+50 PV, plafonné aux PV max).
func UseHealthPotion(character *Character) bool {
	if !inventory.RemoveInventory(&character.Inventory, inventory.ItemHealthPotion) {
		return false
	}
	Heal(character, HealthPotionHeal)
	return true
}

// KnowsSpell indique si le personnage connaît déjà un sort.
func KnowsSpell(character Character, spell spells.Spell) bool {
	for _, s := range character.Spells {
		if s.Name == spell.Name {
			return true
		}
	}
	return false
}

// LearnSpell apprend un sort. Un même sort ne peut pas être appris deux fois.
func LearnSpell(character *Character, spell spells.Spell) bool {
	if KnowsSpell(*character, spell) {
		return false
	}
	character.Spells = append(character.Spells, spell)
	return true
}

// UseSpellbook consomme un "Livre de sort" pour apprendre Boule de Feu.
// Ne consomme rien si le sort est déjà connu.
func UseSpellbook(character *Character) bool {
	if !inventory.HasItem(character.Inventory, inventory.ItemSpellbook) {
		return false
	}
	if KnowsSpell(*character, spells.Fireball) {
		fmt.Println("Vous connaissez déjà", spells.Fireball.Name+".")
		return false
	}

	inventory.RemoveInventory(&character.Inventory, inventory.ItemSpellbook)
	LearnSpell(character, spells.Fireball)
	fmt.Println("Vous apprenez", spells.Fireball.Name+" !")
	return true
}
