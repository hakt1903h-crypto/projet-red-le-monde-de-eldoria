package inventory

import (
	"fmt"
)

type Shop struct {
	Items []Item
}

func InitShop() Shop {
	return Shop{
		Items: []Item{
			{Name: "Potion de vie", Price: 3},
			{Name: "Potion de poison", Price: 6},
			{Name: "Livre de sort", Price: 25},
			{Name: "Fourrure de loup", Price: 4},
			{Name: "Peau de Troll", Price: 7},
			{Name: "Cuir de sanglier", Price: 3},
			{Name: "Plume de corbeau", Price: 1},
		},
	}
}

func AccessShop(shop Shop) {
	fmt.Println("════════════════════════")
	fmt.Println("        MARCHAND")
	fmt.Println("════════════════════════")

	for i, item := range shop.Items {
		fmt.Printf("%d. %s - %d or\n", i+1, item.Name, item.Price)
	}
}

func BuyItem(inv *Inventory, gold *int, item Item) bool {
	if *gold < item.Price {
		return false
	}

	if len(inv.Items) >= MaxInventorySize {
		return false
	}

	*gold -= item.Price
	AddInventory(inv, item)

	return true
}
