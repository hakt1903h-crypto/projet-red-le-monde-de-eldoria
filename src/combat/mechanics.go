package combat

import (
	"eldoria/combat/monsters"
	"eldoria/combat/spells"
	"eldoria/input"
	"eldoria/inventory"
	"eldoria/player"
	"fmt"
)

const (
	PoisonDamagePerTick = 10
	PoisonTicks         = 3
)

// ─── Fonctions pures (testables, sans I/O) ───

// CharacterStarts renvoie true si le personnage commence (initiative supérieure ou égale).
func CharacterStarts(character player.Character, monster monsters.Monster) bool {
	return character.Initiative >= monster.Initiative
}

// MonsterDamage renvoie les dégâts du monstre pour un tour donné.
// Tous les 3 tours, l'attaque est renforcée (x2) — le Gobelin (att. 5) fait donc 5 puis 10.
func MonsterDamage(monster monsters.Monster, turn int) int {
	if turn%3 == 0 {
		return monster.Attack * 2
	}
	return monster.Attack
}

// PlayerAttack applique l'attaque de base du personnage au monstre.
func PlayerAttack(character *player.Character, monster *monsters.Monster) int {
	monster.HP -= character.Attack
	if monster.HP < 0 {
		monster.HP = 0
	}
	return character.Attack
}

// CastSpell lance un sort sur un monstre si le mana est suffisant.
func CastSpell(character *player.Character, spell spells.Spell, monster *monsters.Monster) bool {
	if character.Mana < spell.Mana {
		fmt.Println("Mana insuffisant !")
		return false
	}

	character.Mana -= spell.Mana
	monster.HP -= spell.Damage

	if monster.HP < 0 {
		monster.HP = 0
	}

	fmt.Println(character.Name, "utilise", spell.Name+" !")
	fmt.Println("Dégâts :", spell.Damage)
	fmt.Println("Mana restant :", character.Mana, "/", character.MaxMana)

	return true
}

// UsePoisonPotion lance une potion de poison sur le monstre (10 dégâts x 3).
func UsePoisonPotion(character *player.Character, monster *monsters.Monster) bool {
	if !inventory.RemoveInventory(&character.Inventory, inventory.ItemPoisonPotion) {
		return false
	}

	fmt.Println("Vous lancez une potion de poison !")
	for i := 1; i <= PoisonTicks; i++ {
		monster.HP -= PoisonDamagePerTick
		if monster.HP < 0 {
			monster.HP = 0
		}
		fmt.Printf("  Poison %d/%d : -%d PV (%s : %d/%d)\n",
			i, PoisonTicks, PoisonDamagePerTick, monster.Name, monster.HP, monster.MaxHP)
		if monster.HP <= 0 {
			break
		}
	}
	return true
}

// ApplyVictoryRewards distribue XP, or et butin après une victoire.
func ApplyVictoryRewards(character *player.Character, monster *monsters.Monster) {
	if monster.GoldReward > 0 {
		character.Gold += monster.GoldReward
		fmt.Printf("Vous gagnez %d or.\n", monster.GoldReward)
	}
	if monster.XPReward > 0 {
		fmt.Printf("Vous gagnez %d XP.\n", monster.XPReward)
		player.GainXP(character, monster.XPReward)
	}
	for _, item := range monster.Loot {
		if inventory.AddInventory(&character.Inventory, item) {
			fmt.Println("Butin récupéré :", item.Name)
		} else {
			fmt.Println("Inventaire plein : butin", item.Name, "perdu.")
		}
	}
}

// ─── Boucle de combat interactive ───

// Fight déroule un combat tour par tour. Renvoie true si le personnage gagne.
func Fight(character *player.Character, monster *monsters.Monster) bool {
	fmt.Printf("\n════════════════════════════\n")
	fmt.Printf("  COMBAT : %s vs %s\n", character.Name, monster.Name)
	fmt.Printf("════════════════════════════\n")

	playerTurn := CharacterStarts(*character, *monster)
	monsterTurn := 1

	for character.HP > 0 && monster.HP > 0 {
		fmt.Printf("\n%s : %d/%d PV | %s : %d/%d PV\n",
			character.Name, character.HP, character.MaxHP,
			monster.Name, monster.HP, monster.MaxHP)

		if playerTurn {
			playerCombatTurn(character, monster)
		} else {
			dmg := MonsterDamage(*monster, monsterTurn)
			character.HP -= dmg
			if character.HP < 0 {
				character.HP = 0
			}
			if monsterTurn%3 == 0 {
				fmt.Printf("%s utilise une attaque renforcée ! (-%d PV)\n", monster.Name, dmg)
			} else {
				fmt.Printf("%s attaque ! (-%d PV)\n", monster.Name, dmg)
			}
			monsterTurn++
		}

		playerTurn = !playerTurn
	}

	if character.HP <= 0 {
		fmt.Printf("\nVous avez été vaincu par %s...\n", monster.Name)
		return false
	}

	fmt.Printf("\nVictoire ! %s est vaincu.\n", monster.Name)
	ApplyVictoryRewards(character, monster)
	return true
}

func playerCombatTurn(character *player.Character, monster *monsters.Monster) {
	for {
		fmt.Println("── À vous de jouer ──")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Lancer un sort")
		fmt.Println("3. Boire une potion de vie")
		fmt.Println("4. Lancer une potion de poison")
		fmt.Print("> ")

		switch readLine() {
		case "1":
			dmg := PlayerAttack(character, monster)
			fmt.Printf("Vous attaquez %s (-%d PV).\n", monster.Name, dmg)
			return
		case "2":
			if chooseAndCastSpell(character, monster) {
				return
			}
		case "3":
			if player.UseHealthPotion(character) {
				fmt.Printf("Vous buvez une potion de vie (%d/%d PV).\n", character.HP, character.MaxHP)
				return
			}
			fmt.Println("Aucune potion de vie dans l'inventaire.")
		case "4":
			if UsePoisonPotion(character, monster) {
				return
			}
			fmt.Println("Aucune potion de poison dans l'inventaire.")
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func chooseAndCastSpell(character *player.Character, monster *monsters.Monster) bool {
	if len(character.Spells) == 0 {
		fmt.Println("Vous ne connaissez aucun sort.")
		return false
	}

	fmt.Println("Vos sorts :")
	for i, s := range character.Spells {
		fmt.Printf("  %d. %s (dégâts %d, mana %d)\n", i+1, s.Name, s.Damage, s.Mana)
	}
	fmt.Print("Sort > ")

	idx, err := strconv.Atoi(readLine())
	if err != nil || idx < 1 || idx > len(character.Spells) {
		fmt.Println("Sort invalide.")
		return false
	}

	return CastSpell(character, character.Spells[idx-1], monster)
}

// TrainingFight lance le combat d'entraînement contre le Gobelin (fonctionnalité de test du sujet).
func TrainingFight(character *player.Character) {
	goblin := monsters.InitGoblin()
	if !Fight(character, &goblin) {
		player.Revive(character)
	}
}
