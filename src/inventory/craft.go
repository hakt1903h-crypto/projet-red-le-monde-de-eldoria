package inventory

// Recipe décrit une recette de forge : un résultat, un coût en or et des matériaux.
type Recipe struct {
	Result    string
	GoldCost  int
	Materials map[string]int
}

const CraftGoldCost = 5

// Recipes liste les recettes disponibles à la forge.
var Recipes = []Recipe{
	{
		Result:    EquipHat,
		GoldCost:  CraftGoldCost,
		Materials: map[string]int{ItemCrowFeather: 1, ItemBoarLeather: 1},
	},
	{
		Result:    EquipTunic,
		GoldCost:  CraftGoldCost,
		Materials: map[string]int{ItemWolfFur: 2, ItemTrollHide: 1},
	},
	{
		Result:    EquipBoots,
		GoldCost:  CraftGoldCost,
		Materials: map[string]int{ItemWolfFur: 1, ItemBoarLeather: 1},
	},
}

func FindRecipe(result string) (Recipe, bool) {
	for _, r := range Recipes {
		if r.Result == result {
			return r, true
		}
	}
	return Recipe{}, false
}

// CanCraft vérifie l'or et les matériaux sans rien consommer.
func CanCraft(inv Inventory, gold int, recipe Recipe) bool {
	if gold < recipe.GoldCost {
		return false
	}
	for mat, qty := range recipe.Materials {
		if CountItem(inv, mat) < qty {
			return false
		}
	}
	return true
}

// CraftItem fabrique un équipement : retire l'or et les matériaux, ajoute le résultat.
func CraftItem(inv *Inventory, gold *int, recipe Recipe) bool {
	if !CanCraft(*inv, *gold, recipe) {
		return false
	}

	*gold -= recipe.GoldCost
	for mat, qty := range recipe.Materials {
		for i := 0; i < qty; i++ {
			RemoveInventory(inv, mat)
		}
	}

	AddInventory(inv, Item{Name: recipe.Result, Type: TypeEquipment})
	return true
}
