package models

type ItemType string

const (
	HealthPotion     ItemType = "HEALTH_POTION"
	PoisonPotion     ItemType = "POISON_POTION"
	SpellBook        ItemType = "SPELL_BOOK"
	Resource         ItemType = "RESOURCE"
	EquipmentItem    ItemType = "EQUIPMENT"
	InventoryUpgrade ItemType = "INVENTORY_UPGRADE"
)

type Item struct {
	Name    string
	Type    ItemType
	Price   int
	Effect  int
	Spell   *Spell
	Slot    string
	HPBonus int
}
