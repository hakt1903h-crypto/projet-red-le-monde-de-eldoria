package player

import (
	"eldoria/combat/spells"
	"eldoria/inventory"
	"fmt"
)

type Character struct {
	Name  string
	Class string

	Level int

	MaxHP int
	HP    int

	Inventory inventory.Inventory
	Gold      int

	Spells []spells.Spell
	Attack int

	Mana    int
	MaxMana int

	XP    int
	MaxXP int

	Initiative int

	Equipment map[string]string // slot -> nom de l'équipement porté
}

func InitCharacter() Character {
	return Character{
		Level:      1,
		HP:         50,
		MaxHP:      100,
		Mana:       50,
		MaxMana:    50,
		XP:         0,
		MaxXP:      100,
		Gold:       100,
		Attack:     5,
		Inventory:  inventory.Inventory{},
		Spells:     []spells.Spell{},
		Equipment:  map[string]string{},
		Initiative: 10,
	}
}
func CharacterCreation() Character {
	character := InitCharacter()

	fmt.Println("Quel est votre nom ? ")
	fmt.Scanln(&character.Name)

	fmt.Println()
	fmt.Println("Choisissez votre classe :")
	fmt.Println("1. Guerrier")
	fmt.Println("2. Mage")
	fmt.Println("3. Voleur")

	var choice int

	for {
		fmt.Print("> ")
		fmt.Scanln(&choice)

		if choice >= 1 && choice <= 3 {
			break
		}

		fmt.Println("Choix invalide.")
	}

	switch choice {
	case 1:
		character.Class = "Guerrier"
	case 2:
		character.Class = "Mage"
	case 3:
		character.Class = "Voleur"
	}
	ApplyClass(&character, character.Class)

	return character
}
func ApplyClass(character *Character, class string) {
	switch class {
	case "Guerrier":
		character.MaxHP = 120
		character.MaxMana = 30
		character.Mana = 30
		character.Initiative = 8
		character.Attack = 10
		character.Spells = []spells.Spell{spells.Punch}

	case "Mage":
		character.MaxHP = 80
		character.MaxMana = 100
		character.Mana = 100
		character.Initiative = 10
		character.Attack = 5
		character.Spells = []spells.Spell{spells.Punch, spells.Fireball}

	case "Voleur":
		character.MaxHP = 90
		character.MaxMana = 50
		character.Mana = 50
		character.Initiative = 15
		character.Attack = 8
		character.Spells = []spells.Spell{spells.Punch}
	}

	// PV de départ : 50 % des PV maximum (consigne du sujet)
	character.HP = character.MaxHP / 2
}

func DisplayInfo(character Character) {
	fmt.Println("════════════════════════════════")
	fmt.Println("         PERSONNAGE")
	fmt.Println("════════════════════════════════")

	fmt.Println("Nom :", character.Name)
	fmt.Println("Classe :", character.Class)
	fmt.Println("Niveau :", character.Level)
	fmt.Println()
	fmt.Printf("PV : %d/%d\n", character.HP, character.MaxHP)
	fmt.Printf("Mana : %d/%d\n", character.Mana, character.MaxMana)
	fmt.Printf("XP : %d/%d\n", character.XP, character.MaxXP)
	fmt.Println()
	fmt.Println("Initiative :", character.Initiative)
	fmt.Println("Attaque :", character.Attack)
	fmt.Println("Argent :", character.Gold, "or")

	fmt.Print("Sorts : ")
	if len(character.Spells) == 0 {
		fmt.Println("aucun")
	} else {
		for i, s := range character.Spells {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(s.Name)
		}
		fmt.Println()
	}

	fmt.Print("Équipement : ")
	if len(character.Equipment) == 0 {
		fmt.Println("aucun")
	} else {
		first := true
		for slot, name := range character.Equipment {
			if !first {
				fmt.Print(", ")
			}
			fmt.Printf("%s (%s)", name, slot)
			first = false
		}
		fmt.Println()
	}

	fmt.Println("════════════════════════════════")
}
