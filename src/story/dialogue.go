package story

// Choice est une option de dialogue proposée au joueur, avec la réponse déclenchée.
type Choice struct {
	Text     string
	Response string
}

// Dialogue est une réplique de PNJ réutilisable, avec choix optionnels.
// Volontairement découplé de toute interface : le rendu est laissé à l'appelant.
type Dialogue struct {
	Speaker string
	Lines   []string
	Choices []Choice
}

// Dialogues des PNJ du village de Néris, indexés par nom.
var Dialogues = map[string]Dialogue{
	"Arel": {
		Speaker: "Arel",
		Lines: []string{
			"Je savais que tu reviendrais.",
			"Le temps t'a effacé, mais pas complètement.",
		},
		Choices: []Choice{
			{Text: "Qui êtes-vous ?", Response: "Un vieil archiviste. J'ai consigné ton histoire quand tu l'as oubliée."},
			{Text: "Comment connaissez-vous mon nom ?", Response: "Tout Eldoria connaît ton nom. C'est précisément le problème."},
		},
	},
	"Kael": {
		Speaker: "Kael",
		Lines: []string{
			"Reste où je peux te voir, étranger.",
			"On dit que là où tu passes, le monde s'efface un peu plus.",
		},
		Choices: []Choice{
			{Text: "Je ne veux de mal à personne.", Response: "Les intentions ne nourrissent pas la méfiance. Prouve-le."},
			{Text: "Que sais-tu de moi ?", Response: "Assez pour te surveiller. Pas assez pour te faire confiance."},
		},
	},
	"Mira": {
		Speaker: "Mira",
		Lines: []string{
			"Tiens, bois cette infusion. Tu as l'air épuisé.",
			"Je... je crois que je t'ai déjà soigné, avant. Mais le souvenir me glisse entre les doigts.",
		},
		Choices: []Choice{
			{Text: "Merci, Mira.", Response: "Reviens quand tu veux. Tant que je me souviens encore de toi."},
			{Text: "Tu m'oublies, toi aussi ?", Response: "Un peu plus chaque jour. Fais vite, avant qu'il ne reste rien."},
		},
	},
	"Le Veilleur": {
		Speaker: "Le Veilleur",
		Lines: []string{
			"Chaque fragment que tu retrouves rend au monde une couleur.",
			"Mais souviens-toi : ce que tu as scellé attend aussi de revenir.",
		},
	},
}

// DialogueOf renvoie le dialogue d'un PNJ.
func DialogueOf(npc string) (Dialogue, bool) {
	d, ok := Dialogues[npc]
	return d, ok
}
