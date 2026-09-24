package character

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"

	"eldoria/interne/models"
)

func CreateCharacter(scanner *bufio.Scanner) models.Character {
	fmt.Println("=== Création du personnage ===")

	name := askName(scanner)
	class := askClass(scanner)

	maxHP := getMaxHP(class)

	player := models.Character{
		Name:              name,
		Class:             class,
		Level:             1,
		MaxHP:             maxHP,
		CurrentHP:         maxHP / 2,
		Inventory:         []models.Item{},
		InventoryCapacity: models.DefaultInventorySize,
		Money:             models.StartingGold,
		Mana:              models.DefaultMaxMana,
		MaxMana:           models.DefaultMaxMana,
		Spells: []models.Spell{
			{
				Name:     "Coup de Poing",
				Damage:   models.PunchDamage,
				ManaCost: 0,
			},
		},
	}

	return player
}

func askName(scanner *bufio.Scanner) string {
	for {
		fmt.Print("Nom : ")

		if !scanner.Scan() {
			return "Joueur"
		}

		name, err := NormalizeName(scanner.Text())
		if err != nil {
			fmt.Println("Nom invalide :", err)
			continue
		}

		return name
	}
}

func askClass(scanner *bufio.Scanner) string {
	for {
		fmt.Println()
		fmt.Println("Choisissez votre classe :")
		fmt.Println("1. Humain")
		fmt.Println("2. Elfe")
		fmt.Println("3. Mage")
		fmt.Print("> ")

		if !scanner.Scan() {
			return "Humain"
		}

		switch strings.ToLower(strings.TrimSpace(scanner.Text())) {
		case "1", "humain":
			return "Humain"
		case "2", "elfe":
			return "Elfe"
		case "3", "Mage":
			return "mage"
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func getMaxHP(class string) int {
	switch class {
	case "Humain":
		return models.HumanMaxHP
	case "Elfe":
		return models.ElfMaxHP
	case "Mage":
		return models.MageMaxHP
	default:
		return models.HumanMaxHP
	}
}

func NormalizeName(input string) (string, error) {
	name := strings.TrimSpace(input)

	if name == "" {
		return "", fmt.Errorf("le nom ne peut pas être vide")
	}

	runes := []rune(name)

	for _, r := range runes {
		if !unicode.IsLetter(r) {
			return "", fmt.Errorf("le nom doit contenir uniquement des lettres")
		}
	}

	lowerName := []rune(strings.ToLower(name))
	lowerName[0] = unicode.ToUpper(lowerName[0])

	return string(lowerName), nil
}

func DisplayInfo(player models.Character) {
	fmt.Println()
	fmt.Println("========== INFORMATIONS ==========")
	fmt.Println("Nom :", player.Name)
	fmt.Println("Classe :", player.Class)
	fmt.Println("Niveau :", player.Level)
	fmt.Printf("PV : %d / %d\n", player.CurrentHP, player.MaxHP)
	fmt.Printf("Mana : %d / %d\n", player.Mana, player.MaxMana)
	fmt.Printf("XP : %d / %d\n", player.XP, player.MaxXP)
	fmt.Println("Or :", player.Money)
	fmt.Printf(
		"Inventaire : %d / %d\n",
		len(player.Inventory),
		player.InventoryCapacity,
	)
	fmt.Println("Initiative :", player.Initiative)

	fmt.Println("Sorts :")
	if len(player.Spells) == 0 {
		fmt.Println("  Aucun sort")
	} else {
		for _, spell := range player.Spells {
			fmt.Printf(
				"  - %s (%d dégâts, %d mana)\n",
				spell.Name,
				spell.Damage,
				spell.ManaCost,
			)
		}
	}

	fmt.Println("Équipements :")
	displayEquipmentSlot("Tête", player.Equipment.Head)
	displayEquipmentSlot("Torse", player.Equipment.Chest)
	displayEquipmentSlot("Pieds", player.Equipment.Feet)

	fmt.Println("==================================")
	fmt.Println()
}

func displayEquipmentSlot(slotName string, item *models.Item) {
	if item == nil {
		fmt.Printf("  %s : vide\n", slotName)
		return
	}

	fmt.Printf("  %s : %s\n", slotName, item.Name)
}

func IsDead(player models.Character) bool {
	return player.CurrentHP <= 0
}

func HandleDeath(player *models.Character) {
	if !IsDead(*player) {
		return
	}

	player.CurrentHP = player.MaxHP / 2
}

func AddMoney(player *models.Character, amount int) {
	if amount <= 0 {
		return
	}

	player.Money += amount
}

func RemoveMoney(player *models.Character, amount int) bool {
	if amount < 0 || player.Money < amount {
		return false
	}

	player.Money -= amount
	return true
}

func CanAfford(player models.Character, amount int) bool {
	return amount >= 0 && player.Money >= amount
}

func HasSpell(player models.Character, spellName string) bool {
	for _, spell := range player.Spells {
		if strings.EqualFold(spell.Name, spellName) {
			return true
		}
	}

	return false
}

func LearnSpell(player *models.Character, spell models.Spell) bool {
	if HasSpell(*player, spell.Name) {
		return false
	}

	player.Spells = append(player.Spells, spell)
	return true
}

func UseSpellBook(player *models.Character, itemIndex int) bool {
	if itemIndex < 0 || itemIndex >= len(player.Inventory) {
		return false
	}

	item := player.Inventory[itemIndex]

	if item.Type != models.SpellBook || item.Spell == nil {
		return false
	}

	if !LearnSpell(player, *item.Spell) {
		return false
	}

	player.Inventory = append(
		player.Inventory[:itemIndex],
		player.Inventory[itemIndex+1:]...,
	)

	return true
}

func CastSpell(player *models.Character, spellName string) (int, bool) {
	if player.SilencedTurns > 0 {
		return 0, false
	}

	for _, spell := range player.Spells {
		if !strings.EqualFold(spell.Name, spellName) {
			continue
		}

		if player.Mana < spell.ManaCost {
			return 0, false
		}

		player.Mana -= spell.ManaCost
		return spell.Damage, true
	}

	return 0, false
}

func EquipItem(player *models.Character, itemIndex int) bool {
	if itemIndex < 0 || itemIndex >= len(player.Inventory) {
		return false
	}

	item := player.Inventory[itemIndex]

	if item.Type != models.EquipmentItem {
		return false
	}

	switch item.Slot {
	case "HEAD":
		if player.Equipment.Head != nil {
			oldItem := *player.Equipment.Head

			if len(player.Inventory) >= player.InventoryCapacity {
				return false
			}

			player.Inventory = append(player.Inventory, oldItem)
		}

		player.Equipment.Head = &item

	case "CHEST":
		if player.Equipment.Chest != nil {
			oldItem := *player.Equipment.Chest

			if len(player.Inventory) >= player.InventoryCapacity {
				return false
			}

			player.Inventory = append(player.Inventory, oldItem)
		}

		player.Equipment.Chest = &item

	case "FEET":
		if player.Equipment.Feet != nil {
			oldItem := *player.Equipment.Feet

			if len(player.Inventory) >= player.InventoryCapacity {
				return false
			}

			player.Inventory = append(player.Inventory, oldItem)
		}

		player.Equipment.Feet = &item

	default:
		return false
	}

	player.Inventory = append(
		player.Inventory[:itemIndex],
		player.Inventory[itemIndex+1:]...,
	)

	RecalculateStats(player)

	return true
}

func UnequipItem(player *models.Character, slot string) bool {
	if len(player.Inventory) >= player.InventoryCapacity {
		return false
	}

	switch strings.ToUpper(slot) {
	case "HEAD":
		if player.Equipment.Head == nil {
			return false
		}

		player.Inventory = append(
			player.Inventory,
			*player.Equipment.Head,
		)
		player.Equipment.Head = nil

	case "CHEST":
		if player.Equipment.Chest == nil {
			return false
		}

		player.Inventory = append(
			player.Inventory,
			*player.Equipment.Chest,
		)
		player.Equipment.Chest = nil

	case "FEET":
		if player.Equipment.Feet == nil {
			return false
		}

		player.Inventory = append(
			player.Inventory,
			*player.Equipment.Feet,
		)
		player.Equipment.Feet = nil

	default:
		return false
	}

	RecalculateStats(player)

	return true
}

func RecalculateStats(player *models.Character) {
	baseHP := getMaxHP(player.Class) + player.LevelHPBonus

	bonusHP := 0

	if player.Equipment.Head != nil {
		bonusHP += player.Equipment.Head.HPBonus
	}

	if player.Equipment.Chest != nil {
		bonusHP += player.Equipment.Chest.HPBonus
	}

	if player.Equipment.Feet != nil {
		bonusHP += player.Equipment.Feet.HPBonus
	}

	player.MaxHP = baseHP + bonusHP

	if player.CurrentHP > player.MaxHP {
		player.CurrentHP = player.MaxHP
	}

	player.MaxMana = models.DefaultMaxMana + player.LevelManaBonus

	if player.Mana > player.MaxMana {
		player.Mana = player.MaxMana
	}
}
