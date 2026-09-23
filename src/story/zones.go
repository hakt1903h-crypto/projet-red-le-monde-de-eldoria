package story

// Zone décrit une zone du monde d'Eldoria et sa progression narrative.
type Zone struct {
	ID     int
	Name   string
	Intro  string
	Boss   string // nom du boss de fin de zone ("" si aucune)
	IsTown bool   // vrai pour les zones-villages (marchand / forge / PNJ)
}

// Zones : les 5 zones du jeu, dans l'ordre de progression.
var Zones = []Zone{
	{
		ID: 1, Name: "Forêt Oubliée",
		Intro: "Vous vous réveillez sous les arbres, sans presque aucun souvenir.\n" +
			"Le vent murmure un nom... le vôtre. Comment est-ce possible ?",
		Boss: "Loup Cendré",
	},
	{
		ID: 2, Name: "Village de Néris",
		Intro: "Un village apparaît entre les brumes. Les habitants s'arrêtent et vous fixent.\n" +
			"« Vous êtes revenu... », murmure une voix. Ici vivent Arel, Kael et Mira.",
		IsTown: true,
	},
	{
		ID: 3, Name: "Ruines d'Astra",
		Intro: "Les vestiges d'une civilisation oubliée s'étendent devant vous.\n" +
			"Sur les murs, des fresques... et un visage qui ressemble étrangement au vôtre.",
		Boss: "Gardien Oublié",
	},
	{
		ID: 4, Name: "Temple de la Mémoire",
		Intro: "Le cœur du monde. Ici, la vérité affleure à chaque pas.\n" +
			"L'Oubli rôde, avide de dévorer ce qui reste de vous.",
		Boss: "Dévoreur de Souvenirs",
	},
	{
		ID: 5, Name: "Château d'Eldoria",
		Intro: "Les portes du château se dressent, silencieuses.\n" +
			"Au bout de ce chemin vous attend celui qui a tout oublié : le Roi Sans-Visage.",
		Boss: "Le Roi Sans-Visage",
	},
}

// ZoneByID renvoie la zone d'identifiant donné.
func ZoneByID(id int) (Zone, bool) {
	for _, z := range Zones {
		if z.ID == id {
			return z, true
		}
	}
	return Zone{}, false
}

// ZoneCount renvoie le nombre total de zones.
func ZoneCount() int {
	return len(Zones)
}
