package inventory

import "fmt"

const MaxInventorySize = 10

type Inventory struct {
	Items []Item
}

func AddInventory(inv *Inventory, item Item) bool {
	if len(inv.Items) >= MaxInventorySize {
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
	for _, item := range inv.Items {
		if item.Name == itemName {
			return true
		}
	}

	return false
}
func ShowInventory(inv Inventory) {
	fmt.Println("════════════════════════")
	fmt.Println("       INVENTAIRE")
	fmt.Println("════════════════════════")

	if len(inv.Items) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}

	for i, item := range inv.Items {
		fmt.Printf("%d. %s\n", i+1, item.Name)
	}

	fmt.Printf("\n%d/10 objets\n", len(inv.Items))
}
