package entity

import "fmt"

type Character struct {
	Name  string
	Class string

	Level int

	MaxHP int
	HP    int

	Inventory []string
	Gold      int

	Spells []string

	Mana    int
	MaxMana int

	XP    int
	MaxXP int

	Initiative int

	Equipment []string
}

func InitCharacter() Character {
	return Character{
		Level:      1,
		HP:         100,
		MaxHP:      100,
		Mana:       50,
		MaxMana:    50,
		XP:         0,
		MaxXP:      100,
		Gold:       0,
		Inventory:  []string{},
		Spells:     []string{},
		Equipment:  []string{},
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
		character.HP = 120
		character.MaxMana = 30
		character.Mana = 30
		character.Initiative = 8

	case "Mage":
		character.MaxHP = 80
		character.HP = 80
		character.MaxMana = 100
		character.Mana = 100
		character.Initiative = 10

	case "Voleur":
		character.MaxHP = 90
		character.HP = 90
		character.MaxMana = 50
		character.Mana = 50
		character.Initiative = 15
	}
}
