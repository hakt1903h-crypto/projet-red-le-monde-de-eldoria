package main

import (
	"fmt"
	"time"
)

// ==========================================
// 🧩 OUTILS DE BASE : MANIPULATION DE RUNES (Style Piscine Go)
// ==========================================

// RuneEquals compare deux chaînes caractère par caractère (rune par rune) 
// de manière insensible à la casse, sans utiliser strings.EqualFold.
func RuneEquals(s1, s2 string) bool {
	runes1 := []rune(s1)
	runes2 := []rune(s2)

	if len(runes1) != len(runes2) {
		return false
	}

	for i := 0; i < len(runes1); i++ {
		r1 := runes1[i]
		r2 := runes2[i]

		// Conversion manuelle en minuscule si c'est une majuscule (ASCII A-Z -> a-z)
		if r1 >= 'A' && r1 <= 'Z' {
			r1 += 32
		}
		if r2 >= 'A' && r2 <= 'Z' {
			r2 += 32
		}

		if r1 != r2 {
			return false
		}
	}
	return true
}

// RuneLen retourne la longueur réelle d'une chaîne en comptant ses runes (gère les accents/emojis)
func RuneLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

// ==========================================
// 🎒 PARTIE 1 : GESTION DE L'INVENTAIRE (Tâches 9 & 10)
// ==========================================

// AddItem ajoute un objet à l'inventaire dans la limite de la taille maximale
func AddItem(c *Character, item Item) bool {
	if len(c.Inventory) >= c.MaxInvSize {
		fmt.Println("❌ Inventaire plein ! Impossible d'ajouter l'objet.")
		return false
	}
	c.Inventory = append(c.Inventory, item)
	fmt.Printf("✨ Objet obtenu : %s (%s)\n", item.Name, item.Description)
	return true
}

// RemoveItem retire un objet de l'inventaire en utilisant notre comparaison par runes
func RemoveItem(c *Character, itemName string) (Item, bool) {
	for i, item := range c.Inventory {
		// On utilise notre fonction RuneEquals maison (façon piscine)
		if RuneEquals(item.Name, itemName) {
			// Supprime l'élément du slice en préservant l'ordre
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return item, true
		}
	}
	return Item{}, false
}

// ExpandInventory permet d'agrandir l'inventaire de 10 emplacements (Max 3 fois, Tâche 9)
func ExpandInventory(c *Character) bool {
	const maxExtensions = 3
	currentExtensions := (c.MaxInvSize - 10) / 10

	if currentExtensions >= maxExtensions {
		fmt.Println("❌ Votre inventaire est déjà à sa taille maximale (40 emplacements).")
		return false
	}

	cost := 30 * (currentExtensions + 1)
	if c.Money < cost {
		fmt.Printf("❌ Pas assez d'or ! Il vous faut %d pièces d'or.\n", cost)
		return false
	}

	c.Money -= cost
	c.MaxInvSize += 10
	fmt.Printf("🎒 Extension réussie ! Capacité : %d emplacements (Coût : %d or).\n", c.MaxInvSize, cost)
	return true
}

// ==========================================
// 🧪 PARTIE 2 : POTIONS & EFFETS (Tâche 7)
// ==========================================

// UseHealthPotion soigne le personnage de 50 PV (sans dépasser MaxHP)
func UseHealthPotion(c *Character) {
	item, found := RemoveItem(c, "Élixir de Sang de Lune")
	if !found {
		_, found = RemoveItem(c, "Potion de vie")
		if !found {
			fmt.Println("❌ Vous n'avez pas de Potion de vie dans votre inventaire.")
			return
		}
	}

	healAmount := 50
	oldHP := c.CurrentHP
	c.CurrentHP += healAmount
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}

	fmt.Printf("🧪 %s boit une potion de soin. PV : %d/%d (+%d PV)\n", c.Name, c.CurrentHP, c.MaxHP, c.CurrentHP-oldHP)
	_ = item
}

// UsePoisonPotion inflige 10 dégâts par seconde pendant 3 secondes (Tâche 7)
func UsePoisonPotion(c *Character) {
	item, found := RemoveItem(c, "Fiole de Venin Nécromantique")
	if !found {
		_, found = RemoveItem(c, "Potion de poison")
		if !found {
			fmt.Println("❌ Vous n'avez pas de Potion de poison.")
			return
		}
	}

	fmt.Printf("☠️ %s boit une Fiole de Venin Nécromantique ! Le poison se propage...\n", c.Name)

	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		c.CurrentHP -= 10
		fmt.Printf("   [Sec %d/3] -10 PV subis. (PV actuels : %d/%d)\n", i, c.CurrentHP, c.MaxHP)
		
		if CheckDeath(c) {
			fmt.Println("💀 Le poison vous a achevé...")
			return
		}
	}
	_ = item
}

// ==========================================
// 💀 PARTIE 3 : POINTS DE VIE & MORT (Tâche 8 & 12)
// ==========================================

// CheckDeath vérifie si les PV tombent à 0 et réanime à 50% des PV max
func CheckDeath(c *Character) bool {
	if c.CurrentHP <= 0 {
		c.CurrentHP = c.MaxHP / 2
		fmt.Println("\n╔══════════════════════════════════════════════════╗")
		fmt.Println("║               VOUS ÊTES MORT...                  ║")
		fmt.Println("║     Les ténèbres vous rejettent à la vie.        ║")
		fmt.Printf("║     Ressurrection à %d / %d PV                   ║\n", c.CurrentHP, c.MaxHP)
		fmt.Println("╚══════════════════════════════════════════════════╝")
		return true
	}
	return false
}

// ==========================================
// 💰 PARTIE 4 : MONNAIE & MARCHAND (Tâches 13 & 14)
// ==========================================

// AddMoney ajoute de l'or au personnage
func AddMoney(c *Character, amount int) {
	c.Money += amount
	fmt.Printf("💰 Vous avez gagné %d pièces d'or. (Total : %d or)\n", amount, c.Money)
}

// SpendMoney retire de l'or si le joueur en a assez
func SpendMoney(c *Character, amount int) bool {
	if c.Money < amount {
		fmt.Printf("❌ Fonds insuffisants ! Il vous manque %d pièces d'or.\n", amount-c.Money)
		return false
	}
	c.Money -= amount
	fmt.Printf("💸 Vous avez dépensé %d pièces d'or. (Restant : %d or)\n", amount, c.Money)
	return true
}

// BuyItem permet d'acheter un objet chez le marchand
func BuyItem(c *Character, itemName string, price int, itemType string) {
	if !SpendMoney(c, price) {
		return
	}
	
	newItem := Item{
		Name:        itemName,
		Description: "Acheté auprès du Marchand",
		Type:        itemType,
		Price:       price,
	}

	AddItem(c, newItem)
}

// ==========================================
// 🔨 PARTIE 5 : FORGERON & ÉQUIPEMENT (Tâches 15 à 18)
// ==========================================

// CraftItem permet au forgeron de fabriquer le set d'aventurier
func CraftItem(c *Character, itemName string, requiredMat string, matCostCount int, goldCost int) {
	if c.Money < goldCost {
		fmt.Printf("❌ Pas assez d'or ! Il faut %d pièces d'or.\n", goldCost)
		return
	}

	// Compter les matériaux dans l'inventaire via notre fonction RuneEquals
	foundCount := 0
	for _, item := range c.Inventory {
		if RuneEquals(item.Name, requiredMat) {
			foundCount++
		}
	}

	if foundCount < matCostCount {
		fmt.Printf("❌ Matériaux insuffisants ! Il vous faut %d x '%s' (Vous en avez %d).\n", matCostCount, requiredMat, foundCount)
		return
	}

	// Retirer les matériaux consommés
	for i := 0; i < matCostCount; i++ {
		RemoveItem(c, requiredMat)
	}

	c.Money -= goldCost

	var craftedItem Item
	var bonusHP int

	// On utilise RuneEquals pour identifier l'item à forger
	if RuneEquals(itemName, "Heaume du Veilleur") {
		craftedItem = Item{Name: itemName, Description: "Tête (+10 PV)", Type: "equipement"}
		bonusHP = 10
	} else if RuneEquals(itemName, "Cuirasse des Terres Désolées") {
		craftedItem = Item{Name: itemName, Description: "Torse (+25 PV)", Type: "equipement"}
		bonusHP = 25
	} else if RuneEquals(itemName, "Solerets de l'Ombre") {
		craftedItem = Item{Name: itemName, Description: "Pieds (+15 PV)", Type: "equipement"}
		bonusHP = 15
	} else {
		return
	}

	AddItem(c, craftedItem)
	c.MaxHP += bonusHP
	fmt.Printf("🛠️ Forge réussie ! Vous avez fabriqué %s (+%d PV max).\n", itemName, bonusHP)
}

// EquipItem équipe un objet depuis l'inventaire
func EquipItem(c *Character, itemName string) {
	item, found := RemoveItem(c, itemName)
	if !found {
		fmt.Println("❌ Cet objet n'est pas dans votre inventaire.")
		return
	}

	if RuneEquals(itemName, "Heaume du Veilleur") {
		if c.Equip.Head.Name != "" {
			AddItem(c, c.Equip.Head)
		}
		c.Equip.Head = item
		fmt.Println("🛡️ Vous équipez le Heaume sur votre tête.")
	} else if RuneEquals(itemName, "Cuirasse des Terres Désolées") {
		if c.Equip.Chest.Name != "" {
			AddItem(c, c.Equip.Chest)
		}
		c.Equip.Chest = item
		fmt.Println("🛡️ Vous enfilez la Cuirasse sur votre torse.")
	} else if RuneEquals(itemName, "Solerets de l'Ombre") {
		if c.Equip.Feet.Name != "" {
			AddItem(c, c.Equip.Feet)
		}
		c.Equip.Feet = item
		fmt.Println("🛡️ Vous chaussez les Solerets à vos pieds.")
	} else {
		fmt.Println("❌ Cet objet ne peut pas être équipé.")
		AddItem(c, item)
	}
}