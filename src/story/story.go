package story

// Prologue : texte d'introduction du jeu.
func Prologue() []string {
	return []string{
		"ELDORIA — LES FRAGMENTS DU SOUVENIR",
		"",
		"Vous ouvrez les yeux dans une forêt que vous ne reconnaissez pas.",
		"Vous ne savez ni où vous êtes, ni comment vous êtes arrivé ici.",
		"Une seule certitude vous glace : les habitants, eux, connaissent votre nom.",
	}
}

// BossIntro : réplique scénaristique avant un combat de boss.
func BossIntro(boss string) []string {
	switch boss {
	case "Loup Cendré":
		return []string{
			"Un grondement. Le Loup Cendré surgit des cendres de la forêt.",
			"Ses yeux vous reconnaissent. Vous, non.",
		}
	case "Gardien Oublié":
		return []string{
			"La statue à votre effigie s'anime dans un fracas de pierre.",
			"« Prouve que tu te souviens encore de qui tu étais. »",
		}
	case "Dévoreur de Souvenirs":
		return []string{
			"L'air se déchire. Le Dévoreur de Souvenirs se nourrit de tout ce que vous avez oublié.",
			"Chaque coup porté vous rend une bribe de vérité.",
		}
	case "Le Roi Sans-Visage":
		return []string{
			"Sur le trône d'Eldoria siège une silhouette sans visage.",
			"« Tu m'as créé le jour où tu as choisi d'oublier. »",
		}
	default:
		return []string{"Un ennemi redoutable se dresse devant vous."}
	}
}

// FinalRevelation : la vérité dévoilée après la victoire finale.
func FinalRevelation() []string {
	return []string{
		"Les derniers fragments se rassemblent.",
		"",
		"Vous n'avez pas seulement perdu la mémoire : vous l'aviez sacrifiée.",
		"Autrefois gardien d'Eldoria, vous aviez choisi d'oublier pour contenir l'Oubli.",
		"Mais plus vous oubliiez, plus le monde s'effaçait avec vous.",
		"",
		"En vous souvenant, vous rendez à Eldoria ses couleurs et sa vitalité.",
	}
}

// Morale : la morale finale du jeu.
func Morale() []string {
	return []string{
		"Un monde ne disparaît pas seulement lorsque ses bâtiments tombent.",
		"Il disparaît lorsque plus personne ne se souvient de ce qu'il était.",
		"",
		"Souviens-toi de ce qui compte,",
		"et ton monde redeviendra beau.",
	}
}
