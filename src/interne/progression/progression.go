package progression

import (
	"fmt"

	"eldoria/interne/character"
	"eldoria/interne/models"
)

func GainXP(player *models.Character, amount int) int {
	if amount <= 0 {
		return 0
	}

	if player.MaxXP <= 0 {
		player.MaxXP = models.StartingMaxXP
	}

	player.XP += amount

	levelUps := 0

	for player.XP >= player.MaxXP {
		player.XP -= player.MaxXP
		LevelUp(player)
		levelUps++
	}

	return levelUps
}

func LevelUp(player *models.Character) {
	player.Level++

	player.LevelHPBonus += models.LevelHPIncrease
	player.LevelManaBonus += models.LevelManaIncrease

	character.RecalculateStats(player)

	fmt.Println()
	fmt.Println("=== NIVEAU SUPÉRIEUR ===")
	fmt.Println("Niveau :", player.Level)
	fmt.Println("PV max :", player.MaxHP)
	fmt.Println("Mana max :", player.MaxMana)
	fmt.Println()
}
