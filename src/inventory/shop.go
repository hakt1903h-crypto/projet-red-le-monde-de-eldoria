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
			{Name: ItemHealthPotion, Price: 3, Type: TypeConsumable},
			{Name: ItemPoisonPotion, Price: 6, Type: TypeConsumable},
			{Name: ItemSpellbook, Price: 25, Type: TypeSpellbook},
			{Name: ItemWolfFur, Price: 4, Type: TypeResource},
			{Name: ItemTrollHide, Price: 7, Type: TypeResource},
			{Name: ItemBoarLeather, Price: 3, Type: TypeResource},
			{Name: ItemCrowFeather, Price: 1, Type: TypeResource},
			{Name: ItemInventoryUpgrade, Price: InventoryUpgradeCost, Type: TypeUpgrade},
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

// BuyItem achète un objet standard : vérifie l'or ET la place disponible.
// L'amélioration d'inventaire est un achat spécial géré par UpgradeInventory.
func BuyItem(inv *Inventory, gold *int, item Item) bool {
	if *gold < item.Price {
		return false
	}

	if inv.IsFull() {
		return false
	}

	*gold -= item.Price
	AddInventory(inv, item)

	return true
}
