package story

import "eldoria/player"

// GameState rassemble tout l'état du jeu. C'est la structure unique à sauvegarder.
// Elle ne dépend d'aucun rendu.
type GameState struct {
	Character      player.Character
	CurrentZone    int
	DefeatedBosses map[string]bool
	Fragments      map[int]bool // fragments de mémoire collectés (par ID)
	Scene          Scene
	MusicVolume    int
}

// NewGameState initialise une nouvelle partie à partir d'un personnage créé.
func NewGameState(character player.Character) *GameState {
	return &GameState{
		Character:      character,
		CurrentZone:    1,
		DefeatedBosses: map[string]bool{},
		Fragments:      map[int]bool{},
		Scene:          SceneWorld,
		MusicVolume:    50,
	}
}

// ─── Fragments de mémoire ───

// HasFragment indique si un fragment a déjà été collecté.
func (g *GameState) HasFragment(id int) bool {
	return g.Fragments[id]
}

// CollectFragment enregistre un fragment. Renvoie false s'il était déjà obtenu.
func (g *GameState) CollectFragment(id int) bool {
	if g.Fragments[id] {
		return false
	}
	g.Fragments[id] = true
	return true
}

// CollectedFragments renvoie les fragments collectés, dans l'ordre de définition.
func (g *GameState) CollectedFragments() []Fragment {
	var collected []Fragment
	for _, f := range Fragments {
		if g.Fragments[f.ID] {
			collected = append(collected, f)
		}
	}
	return collected
}

// ─── Boss ───

// BossDefeated indique si un boss a déjà été vaincu.
func (g *GameState) BossDefeated(name string) bool {
	return g.DefeatedBosses[name]
}

// MarkBossDefeated enregistre la défaite d'un boss.
func (g *GameState) MarkBossDefeated(name string) {
	g.DefeatedBosses[name] = true
}

// ─── Options ───

// SetMusicVolume borne le volume de la musique entre 0 et 100.
func (g *GameState) SetMusicVolume(v int) {
	if v < 0 {
		v = 0
	}
	if v > 100 {
		v = 100
	}
	g.MusicVolume = v
}
