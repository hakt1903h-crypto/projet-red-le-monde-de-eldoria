package combat

import (
	"bufio"
	"fmt"
	"strconv"

	"eldoria/interne/character"
	"eldoria/interne/inventory"
	"eldoria/interne/models"
	"eldoria/interne/progression"
)

type Result string

const (
	Victory Result = "VICTORY"
	Defeat  Result = "DEFEAT"
)

func InitGoblin() models.Monster {
	return models.Monster{
		Name:             "Gobelin d'entrainement",
		MaxHP:            40,
		CurrentHP:        40,
		Attack:           5,
		Initiative:       10,
		ExperienceReward: 30,
		Defense:          0,
	}
}

func DetermineFirst(player models.Character, monster models.Monster) bool {
	if player.Initiative >= monster.Initiative {
		return true
	}

	return false
}

func Attack(player models.Character, monster *models.Monster) int {
	damage := models.BasicAttackDamage

	damage -= monster.Defense

	if damage < 0 {
		damage = 0
	}

	monster.CurrentHP -= damage

	if monster.CurrentHP < 0 {
		monster.CurrentHP = 0
	}

	return damage
}

func GoblinPattern(turn int, goblin models.Monster) int {
	damage := goblin.Attack

	if turn%3 == 0 {
		damage *= 2
	}

	return damage
}

func CharacterTurn(
	scanner *bufio.Scanner,
	player *models.Character,
	monster *models.Monster,
) bool {
	for {
		fmt.Println()
		fmt.Println("Votre tour")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Sort")
		fmt.Print("> ")

		if !scanner.Scan() {
			return false
		}

		choice := scanner.Text()

		switch choice {
		case "1":
			damage := Attack(*player, monster)
			fmt.Printf(
				"%s inflige %d dégâts.\n",
				player.Name,
				damage,
			)

			endCharacterTurn(player)
			return true

		case "2":
			if useInventoryInCombat(scanner, player) {
				endCharacterTurn(player)
				return true
			}

		case "3":
			if useSpellInCombat(scanner, player, monster) {
				endCharacterTurn(player)
				return true
			}

		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func useInventoryInCombat(
	scanner *bufio.Scanner,
	player *models.Character,
) bool {
	inventory.AccessInventory(*player)

	if len(player.Inventory) == 0 {
		return false
	}

	fmt.Print("Numéro de l'objet : ")

	if !scanner.Scan() {
		return false
	}

	index, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Numéro invalide.")
		return false
	}

	index--

	if index < 0 || index >= len(player.Inventory) {
		fmt.Println("Objet invalide.")
		return false
	}

	item := player.Inventory[index]

	switch item.Type {
	case models.HealthPotion:
		if inventory.UseHealthPotion(player) {
			fmt.Println("Potion de vie utilisée.")
			return true
		}
	default:
		fmt.Println("Cet objet ne peut pas être utilisé ici.")
	}

	return false
}

func useSpellInCombat(
	scanner *bufio.Scanner,
	player *models.Character,
	monster *models.Monster,
) bool {
	if player.SilencedTurns > 0 {
		fmt.Println("Vous êtes réduit au silence.")
		return false
	}

	if len(player.Spells) == 0 {
		fmt.Println("Vous n'avez aucun sort.")
		return false
	}

	fmt.Println("Sorts disponibles :")

	for i, spell := range player.Spells {
		fmt.Printf(
			"%d. %s - %d dégâts - %d mana\n",
			i+1,
			spell.Name,
			spell.Damage,
			spell.ManaCost,
		)
	}

	fmt.Print("> ")

	if !scanner.Scan() {
		return false
	}

	index, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Choix invalide.")
		return false
	}

	index--

	if index < 0 || index >= len(player.Spells) {
		fmt.Println("Choix invalide.")
		return false
	}

	spell := player.Spells[index]

	damage, ok := character.CastSpell(player, spell.Name)

	if !ok {
		fmt.Println("Impossible de lancer ce sort.")
		return false
	}

	damage -= monster.Defense

	if damage < 0 {
		damage = 0
	}

	monster.CurrentHP -= damage

	if monster.CurrentHP < 0 {
		monster.CurrentHP = 0
	}

	fmt.Printf(
		"%s lance %s et inflige %d dégâts.\n",
		player.Name,
		spell.Name,
		damage,
	)

	return true
}

func endCharacterTurn(player *models.Character) {
	if player.SilencedTurns > 0 {
		player.SilencedTurns--
	}
}

func Fight(
	scanner *bufio.Scanner,
	player *models.Character,
	monster *models.Monster,
) Result {
	fmt.Println()
	fmt.Println("================================")
	fmt.Println("COMBAT :", monster.Name)
	fmt.Println("================================")

	playerFirst := DetermineFirst(*player, *monster)

	if playerFirst {
		fmt.Println("Vous avez l'initiative.")
	} else {
		fmt.Println(monster.Name, "a l'initiative.")
	}

	turn := 1

	for player.CurrentHP > 0 && monster.CurrentHP > 0 {
		fmt.Println()
		fmt.Println("========== TOUR", turn, "==========")
		fmt.Printf(
			"%s : %d/%d PV\n",
			player.Name,
			player.CurrentHP,
			player.MaxHP,
		)
		fmt.Printf(
			"%s : %d/%d PV\n",
			monster.Name,
			monster.CurrentHP,
			monster.MaxHP,
		)

		if playerFirst {
			CharacterTurn(scanner, player, monster)

			if monster.CurrentHP <= 0 {
				break
			}

			monsterTurn(player, monster, turn)
		} else {
			monsterTurn(player, monster, turn)

			if player.CurrentHP <= 0 {
				break
			}

			CharacterTurn(scanner, player, monster)
		}

		turn++
	}

	if player.CurrentHP <= 0 {
		character.HandleDeath(player)
		fmt.Println()
		fmt.Println("Vous avez été vaincu.")
		fmt.Println("Vous revenez à", player.CurrentHP, "PV.")
		return Defeat
	}

	fmt.Println()
	fmt.Println(monster.Name, "est vaincu !")

	levels := progression.GainXP(
		player,
		monster.ExperienceReward,
	)

	fmt.Println(
		"XP gagnée :",
		monster.ExperienceReward,
	)

	if levels > 0 {
		fmt.Println(
			"Nombre de niveaux gagnés :",
			levels,
		)
	}

	return Victory
}

func monsterTurn(
	player *models.Character,
	monster *models.Monster,
	turn int,
) {
	damage := monster.Attack

	if monster.IsBoss {
		damage = BossPattern(*monster, turn)
	} else if monster.Name == "Gobelin d'entrainement" {
		damage = GoblinPattern(turn, *monster)
	}

	fmt.Printf(
		"%s attaque et inflige %d dégâts.\n",
		monster.Name,
		damage,
	)

	player.CurrentHP -= damage

	if player.CurrentHP < 0 {
		player.CurrentHP = 0
	}

	ApplyBossEffect(player, *monster, turn)
}

func InitBlackCrow() models.Monster {
	return models.Monster{
		Name:             "Corbeau Noir",
		MaxHP:            20,
		CurrentHP:        20,
		Attack:           4,
		Initiative:       14,
		ExperienceReward: 10,
	}
}

func InitAshWolf() models.Monster {
	return models.Monster{
		Name:             "Loup Cendré",
		MaxHP:            60,
		CurrentHP:        60,
		Attack:           8,
		Initiative:       12,
		ExperienceReward: 60,
		IsBoss:           true,
		BossType:         "ASH_WOLF",
	}
}

func InitForgottenGuard() models.Monster {
	return models.Monster{
		Name:             "Garde Oublié",
		MaxHP:            100,
		CurrentHP:        100,
		Attack:           7,
		Initiative:       8,
		ExperienceReward: 100,
		Defense:          4,
		IsBoss:           true,
		BossType:         "FORGOTTEN_GUARD",
	}
}

func InitWanderingShadow() models.Monster {
	return models.Monster{
		Name:             "Ombre Errante",
		MaxHP:            35,
		CurrentHP:        35,
		Attack:           6,
		Initiative:       11,
		ExperienceReward: 20,
	}
}

func InitMemoryParasite() models.Monster {
	return models.Monster{
		Name:             "Parasite de Mémoire",
		MaxHP:            50,
		CurrentHP:        50,
		Attack:           7,
		Initiative:       13,
		ExperienceReward: 35,
	}
}

func InitAstraSpecter() models.Monster {
	return models.Monster{
		Name:             "Spectre d'Astra",
		MaxHP:            70,
		CurrentHP:        70,
		Attack:           8,
		Initiative:       10,
		ExperienceReward: 45,
	}
}

func InitMemoryDevourer() models.Monster {
	return models.Monster{
		Name:             "Dévoreur de souvenirs",
		MaxHP:            120,
		CurrentHP:        120,
		Attack:           9,
		Initiative:       14,
		ExperienceReward: 150,
		Defense:          2,
		IsBoss:           true,
		BossType:         "MEMORY_DEVOURER",
	}
}

func InitFacelessKing() models.Monster {
	return models.Monster{
		Name:             "Roi Sans-Visage",
		MaxHP:            180,
		CurrentHP:        180,
		Attack:           12,
		Initiative:       15,
		ExperienceReward: 250,
		Defense:          3,
		IsBoss:           true,
		BossType:         "FACELESS_KING",
	}
}

func BossPattern(monster models.Monster, turn int) int {
	switch monster.BossType {
	case "ASH_WOLF":
		if turn%2 == 0 {
			return monster.Attack * 2
		}

		return monster.Attack

	case "FORGOTTEN_GUARD":
		return monster.Attack

	case "MEMORY_DEVOURER":
		return monster.Attack

	case "FACELESS_KING":
		return monster.Attack

	default:
		return monster.Attack
	}
}

func ApplyBossEffect(
	player *models.Character,
	monster models.Monster,
	turn int,
) {
	if monster.BossType == "MEMORY_DEVOURER" &&
		turn%2 == 0 {
		player.SilencedTurns = 1

		fmt.Println(
			"Le Dévoreur bloque temporairement votre magie.",
		)
	}
}

func TrainingFight(
	scanner *bufio.Scanner,
	player *models.Character,
) Result {
	goblin := InitGoblin()

	fmt.Println()
	fmt.Println("Un Gobelin d'entraînement apparaît !")

	return Fight(
		scanner,
		player,
		&goblin,
	)
}
