package inventory

import (
	"eldoria/interne/models"
	"fmt"
	"strings"
	"time"
)

func AddInventory(player *models.Character, item models.Item) bool {
	if player.InventoryCapacity <= 0 {
		player.InventoryCapacity = models.DefaultInventorySize
	}

	if len(player.Inventory) >= player.InventoryCapacity {
		return false
	}

	player.Inventory = append(player.Inventory, item)
	return true
}

func RemoveInventory(player *models.Character, itemName string) bool {
	for i, item := range player.Inventory {
		if strings.EqualFold(item.Name, itemName) {
			player.Inventory = append(
				player.Inventory[:i],
				player.Inventory[i+1:]...,
			)
			return true
		}
	}

	return false
}

func HasItem(player *models.Character, itemName string) bool {
	for _, item := range player.Inventory {
		if strings.EqualFold(item.Name, itemName) {
			return true
		}
	}

	return false
}

func AccessInventory(player models.Character) {
	fmt.Println()
	fmt.Println("========== INVENTAIRE ==========")

	if len(player.Inventory) == 0 {
		fmt.Println("Inventaire vide.")
	} else {
		for i, item := range player.Inventory {
			fmt.Printf(
				"%d. %s - %d or\n",
				i+1,
				item.Name,
				item.Price,
			)
		}
	}

	fmt.Printf(
		"Places : %d / %d\n",
		len(player.Inventory),
		player.InventoryCapacity,
	)

	fmt.Println("================================")
	fmt.Println()
}

func UseHealthPotion(player *models.Character) bool {
	if !HasItem(player, "Potion de vie") {
		return false
	}

	player.CurrentHP += 50

	if player.CurrentHP > player.MaxHP {
		player.CurrentHP = player.MaxHP
	}

	RemoveInventory(player, "Potion de vie")

	return true
}

func UsePoisonPotion(player *models.Character) bool {
	if !HasItem(player, "Potion de poison") {
		return false
	}

	RemoveInventory(player, "Potion de poison")

	for i := 0; i < 3; i++ {
		player.CurrentHP -= 10

		if player.CurrentHP < 0 {
			player.CurrentHP = 0
		}

		fmt.Printf(
			"Poison : -10 PV (%d / %d)\n",
			player.CurrentHP,
			player.MaxHP,
		)

		if i < 2 {
			time.Sleep(time.Second)
		}
	}

	return true
}

func CountItem(player models.Character, itemName string) int {
	count := 0

	for _, item := range player.Inventory {
		if strings.EqualFold(item.Name, itemName) {
			count++
		}
	}

	return count
}

func UpgradeInventory(player *models.Character) bool {
	const maxUpgrades = 3
	const upgradeCost = 30
	const capacityIncrease = 10

	if player.InventoryUpgrades >= maxUpgrades {
		return false
	}

	if player.Money < upgradeCost {
		return false
	}

	player.Money -= upgradeCost
	player.InventoryCapacity += capacityIncrease
	player.InventoryUpgrades++

	return true
}
