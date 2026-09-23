package monsters

import "eldoria/inventory"

// newMonster construit un monstre avec ses PV pleins.
func newMonster(name string, hp, atk, init, xp, gold int, boss bool, loot ...inventory.Item) Monster {
	return Monster{
		Name:       name,
		MaxHP:      hp,
		HP:         hp,
		Attack:     atk,
		Initiative: init,
		XPReward:   xp,
		GoldReward: gold,
		IsBoss:     boss,
		Loot:       loot,
	}
}

// res crée un objet ressource (butin).
func res(name string) inventory.Item {
	return inventory.Item{Name: name, Type: inventory.TypeResource}
}

// ─── Ennemis communs ───

func NewCorbeauNoir() Monster {
	return newMonster("Corbeau Noir", 25, 6, 12, 20, 5, false, res(inventory.ItemCrowFeather))
}

func NewGardeOublie() Monster {
	return newMonster("Garde Oublié", 45, 9, 8, 35, 12, false, res(inventory.ItemBoarLeather))
}

func NewOmbreErrante() Monster {
	return newMonster("Ombre Errante", 40, 11, 13, 40, 12, false, res(inventory.ItemCrowFeather))
}

func NewParasiteMemoire() Monster {
	return newMonster("Parasite de Mémoire", 55, 12, 11, 50, 15, false, res(inventory.ItemTrollHide))
}

func NewSpectreAstra() Monster {
	return newMonster("Spectre d'Astra", 60, 14, 14, 60, 18, false, res(inventory.ItemWolfFur))
}

// ─── Boss ───

func NewLoupCendre() Monster {
	return newMonster("Loup Cendré", 90, 12, 12, 120, 40, true,
		res(inventory.ItemWolfFur), res(inventory.ItemWolfFur))
}

func NewGardienOublie() Monster {
	return newMonster("Gardien Oublié", 140, 16, 10, 180, 60, true,
		res(inventory.ItemTrollHide), res(inventory.ItemBoarLeather))
}

func NewDevoreurSouvenirs() Monster {
	return newMonster("Dévoreur de Souvenirs", 190, 20, 13, 260, 90, true)
}

func NewRoiSansVisage() Monster {
	return newMonster("Le Roi Sans-Visage", 270, 24, 15, 420, 150, true)
}
