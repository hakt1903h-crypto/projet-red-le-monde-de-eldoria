package game

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"

	"eldoria/interne/character"
	"eldoria/interne/combat"
	"eldoria/interne/crafting"
	"eldoria/interne/inventory"
	"eldoria/interne/models"
	"eldoria/interne/shop"
	"eldoria/interne/story"
)

func NewGame(scanner *bufio.Scanner) models.GameState {
	player := character.CreateCharacter(scanner)

	state := models.GameState{
		Player:            player,
		CurrentZone:       story.Forest,
		CurrentScene:      models.SceneExploration,
		MemoryFragments:   []string{},
		DefeatedBosses:    []string{},
		StoryProgress:     story.ProgressIntro,
		IsGameFinished:    false,
		TrainingCompleted: false,
		VisitedZones:      []string{story.Forest},
	}

	return state
}

func Run(
	scanner *bufio.Scanner,
	state *models.GameState,
) {
	zones := story.NewZones()

	for !state.IsGameFinished {
		currentZone := zones[state.CurrentZone]

		fmt.Println()
		fmt.Println("================================")
		fmt.Println(currentZone.Name)
		fmt.Println("================================")
		fmt.Println(currentZone.Description)

		showMenu(currentZone.Name)

		if !scanner.Scan() {
			return
		}

		choice := scanner.Text()

		switch choice {
		case "1":
			explore(scanner, state)

		case "2":
			move(scanner, state, zones)

		case "3":
			inventory.AccessInventory(state.Player)

		case "4":
			character.DisplayInfo(state.Player)

		case "5":
			talk(scanner, state, currentZone)

		case "6":
			if currentZone.Name == story.Village {
				shopMenu(scanner, state)
			} else {
				ShowArtists()
			}

		case "7":
			if currentZone.Name == story.Village {
				craftMenu(scanner, state)
			} else {
				return
			}

		case "8":
			if currentZone.Name == story.Village {
				ShowArtists()
			} else {
				fmt.Println("Choix invalide.")
			}

		case "9":
			if currentZone.Name == story.Village {
				fmt.Println("Au revoir.")
				return
			}

		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func showMenu(zoneName string) {
	fmt.Println()
	fmt.Println("1. Explorer")
	fmt.Println("2. Se déplacer")
	fmt.Println("3. Inventaire")
	fmt.Println("4. Informations")
	fmt.Println("5. Parler")

	if zoneName == story.Village {
		fmt.Println("6. Marchand")
		fmt.Println("7. Forge")
		fmt.Println("8. Qui sont-ils ?")
		fmt.Println("9. Quitter")
	} else {
		fmt.Println("6. Qui sont-ils ?")
		fmt.Println("7. Quitter")
	}

	fmt.Print("> ")
}

func explore(
	scanner *bufio.Scanner,
	state *models.GameState,
) {
	switch state.CurrentZone {
	case story.Forest:
		exploreForest(scanner, state)

	case story.Village:
		exploreVillage(state)

	case story.Ruins:
		exploreRuins(scanner, state)

	case story.Temple:
		exploreTemple(scanner, state)

	case story.Castle:
		exploreCastle(scanner, state)
	}
}

func exploreForest(
	scanner *bufio.Scanner,
	state *models.GameState,
) {
	if !story.HasMemory(*state, "FRAG-01") {
		story.CollectMemory(state, "FRAG-01")
		fmt.Println("Vous trouvez un fragment de mémoire.")
		return
	}

	if !state.TrainingCompleted {
		state.CurrentScene = models.SceneCombat

		result := combat.TrainingFight(
			scanner,
			&state.Player,
		)

		state.CurrentScene = models.SceneExploration

		if result == combat.Victory {
			state.TrainingCompleted = true
		}

		return
	}

	if !story.IsBossDefeated(
		*state,
		"Loup Cendré",
	) {
		state.CurrentScene = models.SceneCombat

		boss := combat.InitAshWolf()

		result := combat.Fight(
			scanner,
			&state.Player,
			&boss,
		)

		state.CurrentScene = models.SceneExploration

		if result == combat.Victory {
			story.MarkBossDefeated(
				state,
				"Loup Cendré",
			)
		}

		return
	}

	fmt.Println("La forêt ne cache plus aucun danger majeur.")
}

func exploreVillage(state *models.GameState) {
	if !story.HasMemory(*state, "FRAG-02") {
		story.CollectMemory(state, "FRAG-02")
		fmt.Println("Un nouveau souvenir refait surface.")
		return
	}

	fmt.Println("Vous explorez les rues du village.")
}

func exploreRuins(
	scanner *bufio.Scanner,
	state *models.GameState,
) {
	if !story.HasMemory(*state, "FRAG-03") {
		story.CollectMemory(state, "FRAG-03")
		fmt.Println("Vous découvrez une ancienne inscription.")
		return
	}

	if !story.IsBossDefeated(
		*state,
		"Garde Oublié",
	) {
		state.CurrentScene = models.SceneCombat

		boss := combat.InitForgottenGuard()

		result := combat.Fight(
			scanner,
			&state.Player,
			&boss,
		)

		state.CurrentScene = models.SceneExploration

		if result == combat.Victory {
			story.MarkBossDefeated(
				state,
				"Garde Oublié",
			)
		}

		return
	}

	fmt.Println("Les ruines sont silencieuses.")
}

func exploreTemple(
	scanner *bufio.Scanner,
	state *models.GameState,
) {
	if !story.HasMemory(*state, "FRAG-04") {
		story.CollectMemory(state, "FRAG-04")
		fmt.Println("Vous comprenez une partie de votre passé.")
		return
	}

	if !story.IsBossDefeated(
		*state,
		"Dévoreur de souvenirs",
	) {
		state.CurrentScene = models.SceneCombat

		boss := combat.InitMemoryDevourer()

		result := combat.Fight(
			scanner,
			&state.Player,
			&boss,
		)

		state.CurrentScene = models.SceneExploration

		if result == combat.Victory {
			story.MarkBossDefeated(
				state,
				"Dévoreur de souvenirs",
			)
		}

		return
	}

	fmt.Println("Le temple n'a plus rien à cacher.")
}

func exploreCastle(
	scanner *bufio.Scanner,
	state *models.GameState,
) {
	if !story.HasMemory(*state, "FRAG-05") {
		story.CollectMemory(state, "FRAG-05")
		fmt.Println("Les derniers souvenirs vous reviennent.")
		return
	}

	if !story.IsBossDefeated(
		*state,
		"Roi Sans-Visage",
	) {
		state.CurrentScene = models.SceneCombat

		boss := combat.InitFacelessKing()

		result := combat.Fight(
			scanner,
			&state.Player,
			&boss,
		)

		state.CurrentScene = models.SceneExploration

		if result == combat.Victory {
			story.MarkBossDefeated(
				state,
				"Roi Sans-Visage",
			)

			FinishGame(state)
		}

		return
	}
}

func move(
	scanner *bufio.Scanner,
	state *models.GameState,
	zones map[string]models.Zone,
) {
	currentZone := zones[state.CurrentZone]

	directions := make([]string, 0, len(currentZone.Connections))

	for direction := range currentZone.Connections {
		directions = append(directions, direction)
	}

	sort.Strings(directions)

	fmt.Println()

	for i, direction := range directions {
		fmt.Printf(
			"%d. Aller à l'%s\n",
			i+1,
			direction,
		)
	}

	fmt.Print("> ")

	if !scanner.Scan() {
		return
	}

	index, err := strconv.Atoi(scanner.Text())

	if err != nil || index < 1 || index > len(directions) {
		fmt.Println("Destination invalide.")
		return
	}

	direction := directions[index-1]
	destination := currentZone.Connections[direction]

	state.CurrentZone = destination
	state.CurrentScene = models.SceneExploration

	story.UpdateProgressForZone(state)

	if !contains(
		state.VisitedZones,
		destination,
	) {
		state.VisitedZones = append(
			state.VisitedZones,
			destination,
		)
	}

	fmt.Println("Vous arrivez dans :", destination)
}

func talk(
	scanner *bufio.Scanner,
	state *models.GameState,
	zone models.Zone,
) {
	if len(zone.NPCs) == 0 {
		fmt.Println("Personne n'est disponible ici.")
		return
	}

	fmt.Println()
	fmt.Println("PNJ disponibles :")

	for i, npc := range zone.NPCs {
		fmt.Printf(
			"%d. %s\n",
			i+1,
			npc.Name,
		)
	}

	fmt.Print("> ")

	if !scanner.Scan() {
		return
	}

	index, err := strconv.Atoi(scanner.Text())

	if err != nil ||
		index < 1 ||
		index > len(zone.NPCs) {
		fmt.Println("Choix invalide.")
		return
	}

	state.CurrentScene = models.SceneDialogue

	story.TalkToNPC(
		scanner,
		zone.NPCs[index-1],
		state,
	)

	state.CurrentScene = models.SceneExploration
}

func shopMenu(
	scanner *bufio.Scanner,
	state *models.GameState,
) {
	merchant := shop.InitShop()

	for {
		shop.AccessShop(merchant)

		fmt.Println("0. Retour")
		fmt.Print("> ")

		if !scanner.Scan() {
			return
		}

		if scanner.Text() == "0" {
			return
		}

		index, err := strconv.Atoi(scanner.Text())

		if err != nil ||
			index < 1 ||
			index > len(merchant.Items) {
			fmt.Println("Choix invalide.")
			continue
		}

		item := merchant.Items[index-1]

		if shop.BuyItem(
			&state.Player,
			merchant,
			item.Name,
		) {
			fmt.Println("Achat réussi :", item.Name)
		} else {
			fmt.Println("Achat impossible.")
		}
	}
}

func craftMenu(
	scanner *bufio.Scanner,
	state *models.GameState,
) {
	recipes := crafting.GetRecipes()

	for {
		fmt.Println()
		fmt.Println("============= FORGE =============")

		for i, recipe := range recipes {
			fmt.Printf(
				"%d. %s - %d or\n",
				i+1,
				recipe.Name,
				recipe.Cost,
			)
		}

		fmt.Println("0. Retour")
		fmt.Print("> ")

		if !scanner.Scan() {
			return
		}

		if scanner.Text() == "0" {
			return
		}

		index, err := strconv.Atoi(scanner.Text())

		if err != nil ||
			index < 1 ||
			index > len(recipes) {
			fmt.Println("Choix invalide.")
			continue
		}

		recipe := recipes[index-1]

		if crafting.CraftItem(
			&state.Player,
			recipe,
		) {
			fmt.Println("Objet fabriqué :", recipe.Result.Name)
		} else {
			fmt.Println("Fabrication impossible.")
		}
	}
}

func FinishGame(state *models.GameState) {
	state.CurrentScene = models.SceneEnding
	state.StoryProgress = story.ProgressFinal
	state.IsGameFinished = true

	fmt.Println()
	fmt.Println("================================")
	fmt.Println("              FIN")
	fmt.Println("================================")
	fmt.Println()
	fmt.Println(story.EndingText())
	fmt.Println()
}

func ShowArtists() {
	fmt.Println()
	fmt.Println("========== QUI SONT-ILS ? ==========")
	fmt.Println("Artiste caché partie 1 : REMPLACER_PAR_LE_NOM")
	fmt.Println("Artiste caché partie 2 : REMPLACER_PAR_LE_NOM")
	fmt.Println("Artiste caché partie 3 : REMPLACER_PAR_LE_NOM")
	fmt.Println("=====================================")
	fmt.Println()
}

func contains(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}

	return false
}
