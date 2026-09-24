package shop

import (
	"fmt"
	"strings"

	"eldoria/interne/inventory"
	"eldoria/interne/models"
)

type Shop struct {
	Items []models.Item
}

func InitShop() Shop {
	fireball := models.Spell{
		Name:     "Boule de Feu",
		Damage:   models.FireballDamage,
		ManaCost: models.FireballManaCost,
	}

	return Shop{
		Items: []models.Item{
			{
				Name:   "Potion de vie",
				Type:   models.HealthPotion,
				Price:  3,
				Effect: 50,
			},
			{
				Name:   "Potion de poison",
				Type:   models.PoisonPotion,
				Price:  6,
				Effect: 30,
			},
			{
				Name:  "Livre de sort",
				Type:  models.SpellBook,
				Price: 25,
				Spell: &fireball,
			},
			{
				Name:  "Fourrure de Loup",
				Type:  models.Resource,
				Price: 4,
			},
			{
				Name:  "Peau de Troll",
				Type:  models.Resource,
				Price: 7,
			},
			{
				Name:  "Cuir de Sanglier",
				Type:  models.Resource,
				Price: 3,
			},
			{
				Name:  "Plume de Corbeau",
				Type:  models.Resource,
				Price: 1,
			},
		},
	}
}

func AccessShop(shop Shop) {
	fmt.Println()
	fmt.Println("============= MARCHAND =============")

	for i, item := range shop.Items {
		fmt.Printf(
			"%d. %-22s %d or\n",
			i+1,
			item.Name,
			item.Price,
		)
	}

	fmt.Println("=====================================")
	fmt.Println()
}

func BuyItem(
	player *models.Character,
	shop Shop,
	itemName string,
) bool {
	for _, item := range shop.Items {
		if !strings.EqualFold(item.Name, itemName) {
			continue
		}

		if player.Money < item.Price {
			return false
		}

		if len(player.Inventory) >= player.InventoryCapacity {
			return false
		}

		player.Money -= item.Price
		inventory.AddInventory(player, item)

		return true
	}

	return false
}
