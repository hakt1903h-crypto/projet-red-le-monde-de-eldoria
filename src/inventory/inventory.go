package inventory

import "fmt"

const (
	BaseInventorySize    = 10
	InventoryUpgradeSize = 10
	MaxInventoryUpgrades = 3
	InventoryUpgradeCost = 30
)

type Inventory struct {
	Items    []Item
	Upgrades int
}

// Capacity renvoie la capacité courante (base + améliorations achetées).
func (inv Inventory) Capacity() int {
	return BaseInventorySize + inv.Upgrades*InventoryUpgradeSize
}

// IsFull indique si l'inventaire est plein.
func (inv Inventory) IsFull() bool {
	return len(inv.Items) >= inv.Capacity()
}

func AddInventory(inv *Inventory, item Item) bool {
	if inv.IsFull() {
		return false
	}

	inv.Items = append(inv.Items, item)
	return true
}

func RemoveInventory(inv *Inventory, itemName string) bool {
	for i, item := range inv.Items {
		if item.Name == itemName {
			inv.Items = append(inv.Items[:i], inv.Items[i+1:]...)
			return true
		}
	}

	return false
}

func HasItem(inv Inventory, itemName string) bool {
	return CountItem(inv, itemName) > 0
}

// CountItem compte le nombre d'exemplaires d'un objet (utile pour la forge).
func CountItem(inv Inventory, itemName string) int {
	count := 0
	for _, item := range inv.Items {
		if item.Name == itemName {
			count++
		}
	}
	return count
}

// UpgradeInventory achète une extension d'inventaire (+10 places, 30 or, 3 max).
func UpgradeInventory(inv *Inventory, gold *int) bool {
	if inv.Upgrades >= MaxInventoryUpgrades {
		return false
	}
	if *gold < InventoryUpgradeCost {
		return false
	}

	*gold -= InventoryUpgradeCost
	inv.Upgrades++
	return true
}

func ShowInventory(inv Inventory) {
	fmt.Println("════════════════════════")
	fmt.Println("       INVENTAIRE")
	fmt.Println("════════════════════════")

	if len(inv.Items) == 0 {
		fmt.Println("Inventaire vide.")
	} else {
		for i, item := range inv.Items {
			fmt.Printf("%d. %s\n", i+1, item.Name)
		}
	}

	fmt.Printf("\n%d/%d objets\n", len(inv.Items), inv.Capacity())
}
