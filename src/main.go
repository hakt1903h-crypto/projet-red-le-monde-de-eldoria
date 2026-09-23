package main

import (
	"eldoria/interne/models"
	"fmt"
)

func main() {
	player := models.Character{
		Name:      "Lucas",
		Class:     "Humain",
		Level:     1,
		MaxHP:     models.HumanMaxHP,
		CurrentHP: models.HumanMaxHP / 2,
		Money:     models.StartingGold,
	}

	game := models.GameState{
		Player:         player,
		CurrentZone:    "Forêt Oubliée",
		CurrentScene:   "INTRO",
		StoryProgress:  "INTRO",
		IsGameFinished: false,
	}

	fmt.Println("ELDORIA")
	fmt.Println("Les Fragments du Souvenir")
	fmt.Println()

	fmt.Println("Joueur :", game.Player.Name)
	fmt.Println("Classe :", game.Player.Class)
	fmt.Println("PV :", game.Player.CurrentHP, "/", game.Player.MaxHP)
	fmt.Println("Or :", game.Player.Money)
	fmt.Println("Zone :", game.CurrentZone)
}
