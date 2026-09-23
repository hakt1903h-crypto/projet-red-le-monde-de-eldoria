package main

import (
	"eldoria/inventory"
	"eldoria/player"
	"fmt"
)

func main() {
	perso := player.CharacterCreation()
	player.DisplayInfo(perso)
	marchand := inventory.InitShop()
	inventory.AccessShop(marchand)
	popo := inventory.Item{Name: "Potion de vie", Price: 3}
	fmt.Println(perso.Inventory.Items)
	perso.Gold = 50
	if inventory.BuyItem(&perso.Inventory, &perso.Gold, popo) {
		fmt.Println(perso.Inventory.Items)
	} else {
		fmt.Println("Pas assez de tal")
	}
	player.DisplayInfo(perso)
}
