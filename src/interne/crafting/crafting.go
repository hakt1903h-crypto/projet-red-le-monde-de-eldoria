package crafting

import (
	"strings"

	"eldoria/interne/inventory"
	"eldoria/interne/models"
)

type Recipe struct {
	Name        string
	Ingredients map[string]int
	Cost        int
	Result      models.Item
}

func GetRecipes() []Recipe {
	return []Recipe{
		{
			Name: "Chapeau",
			Ingredients: map[string]int{
				"Plume de Corbeau": 1,
				"Cuir de Sanglier": 1,
			},
			Cost: 5,
			Result: models.Item{
				Name:    "Chapeau",
				Type:    models.EquipmentItem,
				Slot:    "HEAD",
				HPBonus: 10,
			},
		},
		{
			Name: "Tunique",
			Ingredients: map[string]int{
				"Fourrure de Loup": 2,
				"Peau de Troll":    1,
			},
			Cost: 5,
			Result: models.Item{
				Name:    "Tunique",
				Type:    models.EquipmentItem,
				Slot:    "CHEST",
				HPBonus: 25,
			},
		},
		{
			Name: "Bottes",
			Ingredients: map[string]int{
				"Fourrure de Loup": 1,
				"Cuir de Sanglier": 1,
			},
			Cost: 5,
			Result: models.Item{
				Name:    "Bottes",
				Type:    models.EquipmentItem,
				Slot:    "FEET",
				HPBonus: 15,
			},
		},
	}
}

func FindRecipe(name string) *Recipe {
	for _, recipe := range GetRecipes() {
		if strings.EqualFold(recipe.Name, name) {
			return &recipe
		}
	}

	return nil
}

func CanCraft(player models.Character, recipe Recipe) bool {
	if player.Money < recipe.Cost {
		return false
	}

	if len(player.Inventory) >= player.InventoryCapacity {
		return false
	}

	for itemName, quantity := range recipe.Ingredients {
		if inventory.CountItem(player, itemName) < quantity {
			return false
		}
	}

	return true
}

func CraftItem(player *models.Character, recipe Recipe) bool {
	if !CanCraft(*player, recipe) {
		return false
	}

	for itemName, quantity := range recipe.Ingredients {
		for i := 0; i < quantity; i++ {
			inventory.RemoveInventory(player, itemName)
		}
	}

	player.Money -= recipe.Cost

	inventory.AddInventory(player, recipe.Result)

	return true
}
