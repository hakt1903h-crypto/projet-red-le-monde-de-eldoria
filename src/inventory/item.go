package inventory

// ItemType catégorise un objet pour déterminer son usage.
type ItemType string

const (
	TypeConsumable ItemType = "consommable"
	TypeResource   ItemType = "ressource"
	TypeEquipment  ItemType = "equipement"
	TypeSpellbook  ItemType = "grimoire"
	TypeUpgrade    ItemType = "amelioration"
)

// Noms canoniques des objets (une seule source de vérité pour éviter les fautes de frappe
// entre marchand, butin des monstres et recettes de forge).
const (
	ItemHealthPotion     = "Potion de vie"
	ItemPoisonPotion     = "Potion de poison"
	ItemSpellbook        = "Livre de sort"
	ItemWolfFur          = "Fourrure de loup"
	ItemTrollHide        = "Peau de Troll"
	ItemBoarLeather      = "Cuir de sanglier"
	ItemCrowFeather      = "Plume de corbeau"
	ItemInventoryUpgrade = "Augmentation d'inventaire"

	EquipHat   = "Chapeau de l'aventurier"
	EquipTunic = "Tunique de l'aventurier"
	EquipBoots = "Bottes de l'aventurier"
)

type Item struct {
	Name  string
	Price int
	Type  ItemType
}
