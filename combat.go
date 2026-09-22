package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Monster struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attack    int
}

func InitGoblin() Monster {
	return Monster{Name: "Gobelin d'entrainement", MaxHP: 40, CurrentHP: 40, Attack: 5}
}

func GoblinPattern(m *Monster, c *Character, turn int) {
	dmg := m.Attack
	if turn%3 == 0 {
		dmg *= 2 // Bonus tous les 3 tours
	}
	c.CurrentHP -= dmg
	if c.CurrentHP < 0 {
		c.CurrentHP = 0
	}
	fmt.Printf("⚔️ %s inflige %d dégâts à %s (PV : %d/%d)\n", m.Name, dmg, c.Name, c.CurrentHP, c.MaxHP)
	CheckDeath(c)
}

func TrainingFight(c *Character) {
	goblin := InitGoblin()
	reader := bufio.NewReader(os.Stdin)
	turn := 1

	fmt.Printf("\n--- COMBAT : %s VS %s ---\n", c.Name, goblin.Name)
	for goblin.CurrentHP > 0 && c.CurrentHP > 0 {
		fmt.Printf("\n[Tour %d] \n1. Attaquer (-5 PV)\n2. Boire Potion\nChoix : ", turn)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "1" {
			goblin.CurrentHP -= 5
			if goblin.CurrentHP < 0 {
				goblin.CurrentHP = 0
			}
			fmt.Printf("✨ Vous infligez 5 dégâts (%s PV : %d/%d)\n", goblin.Name, goblin.CurrentHP, goblin.MaxHP)
		} else if input == "2" {
			UseHealthPotion(c)
		} else {
			continue
		}

		if goblin.CurrentHP <= 0 {
			fmt.Println("🎉 Victoire ! Monstre vaincu.")
			AddMoney(c, 25)
			break
		}

		GoblinPattern(&goblin, c, turn)
		if c.CurrentHP <= 0 {
			fmt.Println("💀 Vous avez perdu le combat...")
			break
		}
		turn++
	}
}
