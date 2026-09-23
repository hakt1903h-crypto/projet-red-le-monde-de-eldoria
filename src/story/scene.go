package story

// Scene représente l'état logique courant du jeu (indépendant du rendu).
// Le backend sait ainsi dans quelle scène il se trouve ; une future TUI pourra s'y brancher.
type Scene int

const (
	SceneMenu Scene = iota
	SceneWorld
	SceneExploration
	SceneDialogue
	SceneCombat
	SceneInventory
	SceneShop
	SceneForge
	ScenePause
	SceneGameOver
	SceneEnd
)

func (s Scene) String() string {
	switch s {
	case SceneMenu:
		return "Menu principal"
	case SceneWorld:
		return "Monde"
	case SceneExploration:
		return "Exploration"
	case SceneDialogue:
		return "Dialogue"
	case SceneCombat:
		return "Combat"
	case SceneInventory:
		return "Inventaire"
	case SceneShop:
		return "Marchand"
	case SceneForge:
		return "Forge"
	case ScenePause:
		return "Pause"
	case SceneGameOver:
		return "Game Over"
	case SceneEnd:
		return "Fin"
	default:
		return "Inconnu"
	}
}
