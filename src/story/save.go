package story

import (
	"encoding/json"
	"os"
)

const DefaultSavePath = "eldoria_save.json"

// Save écrit l'état complet du jeu au format JSON (indépendant du rendu).
func (g *GameState) Save(path string) error {
	data, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// Load recharge un état de jeu depuis un fichier JSON.
func Load(path string) (*GameState, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var g GameState
	if err := json.Unmarshal(data, &g); err != nil {
		return nil, err
	}

	// Sécurité : garantir des maps non nil après chargement.
	if g.DefeatedBosses == nil {
		g.DefeatedBosses = map[string]bool{}
	}
	if g.Fragments == nil {
		g.Fragments = map[int]bool{}
	}
	if g.Character.Equipment == nil {
		g.Character.Equipment = map[string]string{}
	}

	return &g, nil
}

// SaveExists indique si un fichier de sauvegarde est présent.
func SaveExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
