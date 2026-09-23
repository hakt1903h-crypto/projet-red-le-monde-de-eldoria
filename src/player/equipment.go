package player

import "eldoria/inventory"

// Equip équipe un objet depuis l'inventaire.
// Si le slot est déjà occupé : l'ancien équipement est retiré, remis dans l'inventaire,
// le nouveau est équipé et les PV max sont recalculés.
func Equip(character *Character, itemName string) bool {
	eq, ok := inventory.Equipments[itemName]
	if !ok {
		return false // ce n'est pas un équipement
	}
	if !inventory.HasItem(character.Inventory, itemName) {
		return false // pas dans l'inventaire
	}

	// On sort le nouvel équipement de l'inventaire (libère une place).
	inventory.RemoveInventory(&character.Inventory, itemName)

	// Si le slot est occupé, on remet l'ancien dans l'inventaire (place libérée juste avant).
	if _, occupied := character.Equipment[eq.Slot]; occupied {
		Unequip(character, eq.Slot)
	}

	character.Equipment[eq.Slot] = itemName
	character.MaxHP += eq.HPBonus
	return true
}

// Unequip retire l'équipement d'un slot, le remet dans l'inventaire et recalcule les PV max.
func Unequip(character *Character, slot string) bool {
	name, occupied := character.Equipment[slot]
	if !occupied {
		return false
	}
	if character.Inventory.IsFull() {
		return false // pas de place pour ranger l'équipement retiré
	}

	eq := inventory.Equipments[name]
	character.MaxHP -= eq.HPBonus
	if character.HP > character.MaxHP {
		character.HP = character.MaxHP
	}

	delete(character.Equipment, slot)
	inventory.AddInventory(&character.Inventory, inventory.Item{Name: name, Type: inventory.TypeEquipment})
	return true
}
