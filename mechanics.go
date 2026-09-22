package main

import (
	"fmt"
	"time"
)

// Comparaison bas niveau par runes (Style Piscine Go)
func RuneEquals(s1, s2 string) bool {
	r1, r2 := []rune(s1), []rune(s2)
	if len(r1) != len(r2) {
		return false
	}
	for i := range r1 {
		c1, c2 := r1[i], r2[i]
		if c1 >= 'A' && c1 <= 'Z' {
			c1 += 32
		}
		if c2 >= 'A' && c2 <= 'Z' {
			c2 += 32
		}
		if c1 != c2 {
			return false
		}
	}
	return true
}

func AddItem(c *Character, item Item) bool {
	if len(c.Inventory) >= c.MaxInvSize {
		fmt.Println("❌ Inventaire plein !")
		return false
	}
	c.Inventory = append(c.Inventory, item)
	fmt.Printf("✨ Objet obtenu : %s\n", item.Name)
	return true
}

func RemoveItem(c *Character, name string) (Item, bool) {
	for i, item := range c.Inventory {
		if RuneEquals(item.Name, name) {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return item, true
		}
	}
	return Item{}, false
}

func ExpandInventory(c *Character) {
	ext := (c.MaxInvSize - 10) / 10
	if ext >= 3 {
		fmt.Println("❌ Capacité maximale atteinte (40 slots).")
		return
	}
	cost := 30 * (ext + 1)
	if c.Money < cost {
		fmt.Printf("❌ Pas assez d'or (requis : %d).\n", cost)
		return
	}
	c.Money -= cost
	c.MaxInvSize += 10
	fmt.Printf("🎒 Inventaire étendu à %d slots !\n", c.MaxInvSize)
}

func UseHealthPotion(c *Character) {
	if _, found := RemoveItem(c, "Élixir de Sang de Lune"); !found {
		if _, found = RemoveItem(c, "Potion de vie"); !found {
			fmt.Println("❌ Aucune potion de vie dans l'inventaire.")
			return
		}
	}
	c.CurrentHP += 50
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}
	fmt.Printf("🧪 Potion bue ! PV actuels : %d / %d\n", c.CurrentHP, c.MaxHP)
}

func UsePoisonPotion(c *Character) {
	if _, found := RemoveItem(c, "Fiole de Venin Nécromantique"); !found {
		if _, found = RemoveItem(c, "Potion de poison"); !found {
			fmt.Println("❌ Aucune potion de poison.")
			return
		}
	}
	fmt.Println("☠️ Poison ingéré ! Dégâts progressifs sur 3 secondes...")
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		c.CurrentHP -= 10
		fmt.Printf("   [Sec %d/3] -10 PV (PV : %d/%d)\n", i, c.CurrentHP, c.MaxHP)
		if CheckDeath(c) {
			return
		}
	}
}

func CheckDeath(c *Character) bool {
	if c.CurrentHP <= 0 {
		c.CurrentHP = c.MaxHP / 2
		fmt.Printf("💀 VOUS ÊTES MORT... Réanimation à %d / %d PV.\n", c.CurrentHP, c.MaxHP)
		return true
	}
	return false
}

func AddMoney(c *Character, amt int) {
	c.Money += amt
	fmt.Printf("💰 +%d pièces d'or (Total : %d)\n", amt, c.Money)
}

func SpendMoney(c *Character, amt int) bool {
	if c.Money < amt {
		fmt.Println("❌ Fonds insuffisants.")
		return false
	}
	c.Money -= amt
	return true
}

func BuyItem(c *Character, name string, price int, itype string) {
	if SpendMoney(c, price) {
		AddItem(c, Item{Name: name, Description: "Acheté au marchand", Type: itype, Price: price})
	}
}

func CraftItem(c *Character, itemName, mat string, matCost, goldCost int) {
	if !SpendMoney(c, goldCost) {
		return
	}
	count := 0
	for _, item := range c.Inventory {
		if RuneEquals(item.Name, mat) {
			count++
		}
	}
	if count < matCost {
		fmt.Println("❌ Matériaux insuffisants dans l'inventaire.")
		c.Money += goldCost // Remboursement
		return
	}
	for i := 0; i < matCost; i++ {
		RemoveItem(c, mat)
	}

	bonus := 0
	if RuneEquals(itemName, "Heaume du Veilleur") {
		bonus = 10
	} else if RuneEquals(itemName, "Cuirasse des Terres Désolées") {
		bonus = 25
	} else if RuneEquals(itemName, "Solerets de l'Ombre") {
		bonus = 15
	}

	AddItem(c, Item{Name: itemName, Type: "equipement"})
	c.MaxHP += bonus
	fmt.Printf("🛠️ Forge réussie : %s (+%d PV max)\n", itemName, bonus)
}
