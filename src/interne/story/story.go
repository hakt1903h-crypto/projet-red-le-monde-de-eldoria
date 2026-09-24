package story

import (
	"bufio"
	"fmt"
	"strconv"

	"eldoria/interne/combat"
	"eldoria/interne/models"
)

const (
	Forest  = "Forêt Oubliée"
	Village = "Village de Néris"
	Ruins   = "Ruines d'Astra"
	Temple  = "Temple de la Mémoire"
	Castle  = "Château d'Eldoria"

	ProgressIntro   = "INTRO"
	ProgressForest  = "FOREST"
	ProgressVillage = "VILLAGE"
	ProgressRuins   = "RUINS"
	ProgressTemple  = "TEMPLE"
	ProgressFinal   = "FINAL"
)

func NewZones() map[string]models.Zone {
	return map[string]models.Zone{
		Forest: {
			Name:        Forest,
			Description: "Une forêt silencieuse où les souvenirs semblent avoir disparu.",
			Enemies: []models.Monster{
				combat.InitBlackCrow(),
				combat.InitAshWolf(),
			},
			Connections: map[string]string{
				"est": Village,
			},
		},

		Village: {
			Name:        Village,
			Description: "Un village où plusieurs personnes semblent connaître votre histoire.",
			NPCs:        NewVillageNPCs(),
			Connections: map[string]string{
				"ouest": Forest,
				"est":   Ruins,
			},
		},

		Ruins: {
			Name:        Ruins,
			Description: "Les ruines d'une ancienne civilisation.",
			Enemies: []models.Monster{
				combat.InitForgottenGuard(),
				combat.InitWanderingShadow(),
			},
			Connections: map[string]string{
				"ouest": Village,
				"est":   Temple,
			},
		},

		Temple: {
			Name:        Temple,
			Description: "Un temple rempli de fragments liés à votre passé.",
			Enemies: []models.Monster{
				combat.InitMemoryParasite(),
				combat.InitAstraSpecter(),
			},
			NPCs: []models.NPC{
				NewWatcher(),
			},
			Connections: map[string]string{
				"ouest": Ruins,
				"est":   Castle,
			},
		},

		Castle: {
			Name:        Castle,
			Description: "Le château où semblent converger tous vos souvenirs.",
			Connections: map[string]string{
				"ouest": Temple,
			},
		},
	}
}

func GetMemoryFragments() []models.MemoryFragment {
	return []models.MemoryFragment{
		{
			ID:    "FRAG-01",
			Title: "Une voix",
			Text:  "Une voix vous appelle. Vous entendez votre nom. Puis tout devient noir.",
		},
		{
			ID:    "FRAG-02",
			Title: "Le village",
			Text:  "Des visages familiers semblent vous reconnaître.",
		},
		{
			ID:    "FRAG-03",
			Title: "L'inscription",
			Text:  "Une ancienne inscription semble avoir été écrite pour vous.",
		},
		{
			ID:    "FRAG-04",
			Title: "La vérité",
			Text:  "Une partie de votre passé commence à refaire surface.",
		},
		{
			ID:    "FRAG-05",
			Title: "Les derniers souvenirs",
			Text:  "Les derniers morceaux de votre mémoire convergent vers le château.",
		},
	}
}

func CollectMemory(
	state *models.GameState,
	memoryID string,
) bool {
	if HasMemory(*state, memoryID) {
		return false
	}

	state.MemoryFragments = append(
		state.MemoryFragments,
		memoryID,
	)

	state.Player.MemoryFragments = append(
		state.Player.MemoryFragments,
		memoryID,
	)

	return true
}

func HasMemory(
	state models.GameState,
	memoryID string,
) bool {
	for _, id := range state.MemoryFragments {
		if id == memoryID {
			return true
		}
	}

	return false
}

func GetMemories(
	state models.GameState,
) []models.MemoryFragment {
	all := GetMemoryFragments()
	result := []models.MemoryFragment{}

	for _, memory := range all {
		if HasMemory(state, memory.ID) {
			result = append(result, memory)
		}
	}

	return result
}

func MarkBossDefeated(
	state *models.GameState,
	bossName string,
) bool {
	for _, name := range state.DefeatedBosses {
		if name == bossName {
			return false
		}
	}

	state.DefeatedBosses = append(
		state.DefeatedBosses,
		bossName,
	)

	return true
}

func IsBossDefeated(
	state models.GameState,
	bossName string,
) bool {
	for _, name := range state.DefeatedBosses {
		if name == bossName {
			return true
		}
	}

	return false
}

func NewVillageNPCs() []models.NPC {
	return []models.NPC{
		{
			Name: "Arel",
			Dialogues: []models.Dialogue{
				{
					ID:   "start",
					Text: "Je savais que tu reviendrais.",
					Choices: []models.DialogueChoice{
						{
							Text:           "Qui êtes-vous ?",
							NextDialogueID: "who",
						},
						{
							Text:           "Comment connaissez-vous mon nom ?",
							NextDialogueID: "name",
						},
					},
				},
				{
					ID:   "who",
					Text: "Quelqu'un qui se souvient de ce que toi tu as oublié.",
				},
				{
					ID:   "name",
					Text: "Parce que ton histoire a commencé bien avant ton réveil.",
				},
			},
		},

		{
			Name: "Kael",
			Dialogues: []models.Dialogue{
				{
					ID:   "start",
					Text: "La forge peut encore réparer certaines choses.",
				},
			},
		},

		{
			Name: "Mira",
			Dialogues: []models.Dialogue{
				{
					ID:   "start",
					Text: "Certains souvenirs font plus mal que les blessures.",
				},
			},
		},
	}
}

func NewWatcher() models.NPC {
	return models.NPC{
		Name: "Le Veilleur",
		Dialogues: []models.Dialogue{
			{
				ID:   "start",
				Text: "Tu approches de la vérité.",
			},
		},
	}
}

func TalkToNPC(
	scanner *bufio.Scanner,
	npc models.NPC,
	state *models.GameState,
) {
	if len(npc.Dialogues) == 0 {
		fmt.Println("Ce personnage n'a rien à dire.")
		return
	}

	currentID := npc.Dialogues[0].ID

	for {
		dialogue, ok := findDialogue(npc, currentID)

		if !ok {
			return
		}

		fmt.Println()
		fmt.Println(npc.Name + " :")
		fmt.Println(dialogue.Text)

		if len(dialogue.Choices) == 0 {
			return
		}

		for i, choice := range dialogue.Choices {
			fmt.Printf("%d. %s\n", i+1, choice.Text)
		}

		fmt.Print("> ")

		if !scanner.Scan() {
			return
		}

		index, err := strconv.Atoi(scanner.Text())
		if err != nil || index < 1 || index > len(dialogue.Choices) {
			fmt.Println("Choix invalide.")
			continue
		}

		choice := dialogue.Choices[index-1]

		if choice.SetStoryProgress != "" {
			state.StoryProgress = choice.SetStoryProgress
		}

		if choice.MemoryID != "" {
			CollectMemory(state, choice.MemoryID)
		}

		if choice.NextDialogueID == "" {
			return
		}

		currentID = choice.NextDialogueID
	}
}

func findDialogue(
	npc models.NPC,
	id string,
) (models.Dialogue, bool) {
	for _, dialogue := range npc.Dialogues {
		if dialogue.ID == id {
			return dialogue, true
		}
	}

	return models.Dialogue{}, false
}

func UpdateProgressForZone(state *models.GameState) {
	switch state.CurrentZone {
	case Forest:
		state.StoryProgress = ProgressForest
	case Village:
		state.StoryProgress = ProgressVillage
	case Ruins:
		state.StoryProgress = ProgressRuins
	case Temple:
		state.StoryProgress = ProgressTemple
	case Castle:
		state.StoryProgress = ProgressFinal
	}
}

func EndingText() string {
	return "Tes souvenirs sont enfin réunis. L'histoire d'Eldoria peut maintenant révéler la vérité."
}
