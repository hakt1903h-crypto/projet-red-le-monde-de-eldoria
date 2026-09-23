package inventory

// Slots d'équipement disponibles.
const (
	SlotHead = "tête"
	SlotBody = "torse"
	SlotFeet = "pieds"
)

// Equipment décrit une pièce d'équipement : son slot et son bonus de PV max.
type Equipment struct {
	Name    string
	Slot    string
	HPBonus int
}

// Equipments répertorie les équipements et leurs bonus (consigne : forge).
var Equipments = map[string]Equipment{
	EquipHat:   {Name: EquipHat, Slot: SlotHead, HPBonus: 10},
	EquipTunic: {Name: EquipTunic, Slot: SlotBody, HPBonus: 25},
	EquipBoots: {Name: EquipBoots, Slot: SlotFeet, HPBonus: 15},
}

// IsEquipment indique si un objet est une pièce d'équipement équipable.
func IsEquipment(name string) bool {
	_, ok := Equipments[name]
	return ok
}
