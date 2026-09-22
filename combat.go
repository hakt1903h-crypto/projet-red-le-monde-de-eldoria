package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ==========================================
// 👹 TÂCHE 19 : STRUCTURE MONSTRE ET GOBELIN
// ==========================================

// Monster représente un monstre ou le boss d'entraînement
type Monster struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attack    int
}

// InitGoblin initialise les paramètres de base du gobelin d'entraînement (Tâche 19)
func InitGoblin() Monster {
	return Monster{
		Name:      "Gobelin d'entrainement",
		MaxHP:     40,
		CurrentHP: 40,
		Attack:    5,
	}
}

// ==========================================
// 🤖 TÂCHE 20 : PATTERN DE L'INTELLIGENCE ARTIFICIELLE
// ==========================================

// GoblinPattern gère le pattern d'attaque du gobelin selon le tour (Tâche 20)
// Chaque tour, il inflige 100% de son attaque. Tous les 3 tours (3, 6, 9...), il inflige 200%.
func GoblinPattern(m *Monster, c *Character, turnNumber int) {
	damage := m.Attack
	if turnNumber%3 == 0 {
		damage = m.Attack * 2
	}

	c.CurrentHP -= damage
	if c.CurrentHP < 0 {
		c.CurrentHP = 0
	}

	// Affichage exigé par le sujet
	fmt.Printf("⚔️ %s inflige à %s %d de dégâts\n", m.Name, c.Name, damage)
	fmt.Printf("   [%s] « PV : %d / %d »\n", c.Name, c.CurrentHP, c.MaxHP)

	// Vérifier si le joueur est mort et le réanimer si besoin (Tâche 8)
	CheckDeath(c)
}

// ==========================================
// ⚔️ TÂCHE 21 : TOUR DU JOUEUR (ATTAQUE & INVENTAIRE)
// ==========================================

// CharacterTurn simule le tour de jeu du joueur
func CharacterTurn(c *Character, m *Monster) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- VOS ACTIONS DE COMBAT ---")
		fmt.Println("1. Attaquer (Attaque basique : -5 PV au monstre)")
		fmt.Println("2. Inventaire (Utiliser une potion)")
		fmt.Print("Votre choix : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			// Attaque basique : inflige 5 dégâts au monstre (Tâche 21)
			damage := 5
			m.CurrentHP -= damage
			if m.CurrentHP < 0 {
				m.CurrentHP = 0
			}
			fmt.Printf("✨ %s inflige %d dégâts à %s\n", c.Name, damage, m.Name)
			fmt.Printf("   [%s] « PV : %d / %d »\n", m.Name, m.CurrentHP, m.MaxHP)
			return true // Fin du tour du joueur, on passe au monstre

		case "2":
			// Utilisation d'un objet de l'inventaire en plein combat (Tâche 21 & 7)
			if len(c.Inventory) == 0 {
				fmt.Println("❌ Votre inventaire est vide !")
				continue
			}

			fmt.Println("\n--- INVENTAIRE DE COMBAT ---")
			for i, item := range c.Inventory {
				fmt.Printf("%d. %s (%s)\n", i+1, item.Name, item.Description)
			}
			fmt.Println("0. Retour")
			fmt.Print("Choisissez un objet à utiliser : ")

			// On utilise ta fonction de soin ou de poison selon ce que le joueur possède
			// Pour simplifier, on applique une potion de vie si disponible :
			UseHealthPotion(c)
			return true // Après l'objet, c'est au tour du monstre

		default:
			fmt.Println("❌ Choix invalide, veuillez entrer 1 ou 2.")
		}
	}
}

// ==========================================
// 🏟️ TÂCHE 22 : BOUCLE GLOBALE DU COMBAT
// ==========================================

// TrainingFight lance le combat d'entraînement tour par tour (Tâche 22)
func TrainingFight(c *Character) {
	goblin := InitGoblin()
	turnNumber := 1

	fmt.Println("\n╔══════════════════════════════════════════════════╗")
	fmt.Println("║           DÉBUT DU COMBAT D'ENTRAÎNEMENT         ║")
	fmt.Printf("║     %s (PV: %d) VS %s (PV: %d)     ║\n", c.Name, c.CurrentHP, goblin.Name, goblin.CurrentHP)
	fmt.Println("╚══════════════════════════════════════════════════╝")

	for {
		fmt.Printf("\n================ TOUR DE COMBAT N°%d ===============-\n", turnNumber)

		// 1. Tour du joueur (Attaque ou Potion)
		if !CharacterTurn(c, &goblin) {
			continue
		}

		// Vérifier si le monstre est mort
		if goblin.CurrentHP <= 0 {
			fmt.Printf("\n🎉 Victoire ! %s a été vaincu par %s !\n", goblin.Name, c.Name)
			AddMoney(c, 25) // Petite récompense en or pour la victoire
			break
		}

		// 2. Tour du monstre (Pattern du Gobelin)
		GoblinPattern(&goblin, c, turnNumber)

		// Vérifier si le joueur a perdu tous ses PV
		if c.CurrentHP <= 0 {
			fmt.Printf("\n💀 Vous avez perdu le combat face à %s...\n", goblin.Name)
			break
		}

		turnNumber++
	}
}