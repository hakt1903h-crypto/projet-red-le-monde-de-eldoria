package story

// Fragment est un fragment de mémoire : il débloque une information narrative.
type Fragment struct {
	ID    int
	Zone  int
	Title string
	Text  string
}

// Fragments : les données brutes des souvenirs (séparées de tout rendu).
var Fragments = []Fragment{
	{
		ID: 1, Zone: 1, Title: "Un rire dans la forêt",
		Text: "Une forêt.\nDeux personnes rient.\nUne voix prononce votre nom.\nPuis tout devient noir.",
	},
	{
		ID: 2, Zone: 2, Title: "Le nom oublié",
		Text: "Un village en fête.\nDes visages qui vous saluent.\nVous connaissiez chacun d'eux.\nPourquoi les avez-vous quittés ?",
	},
	{
		ID: 3, Zone: 3, Title: "Le visage de pierre",
		Text: "Des ruines immenses.\nUne statue à votre effigie.\nUne inscription : « Celui qui se souvient nous protège. »",
	},
	{
		ID: 4, Zone: 4, Title: "Le sacrifice",
		Text: "Une salle de lumière.\nVous tendez la main vers l'Oubli.\nVous choisissez d'oublier pour contenir la faille.\nVous saviez le prix.",
	},
	{
		ID: 5, Zone: 5, Title: "Le lien",
		Text: "Votre mémoire et celle d'Eldoria ne font qu'une.\nEn vous souvenant, vous rendez au monde ses couleurs.",
	},
}

// FragmentByID renvoie le fragment d'identifiant donné.
func FragmentByID(id int) (Fragment, bool) {
	for _, f := range Fragments {
		if f.ID == id {
			return f, true
		}
	}
	return Fragment{}, false
}

// FragmentByZone renvoie le fragment associé à une zone.
func FragmentByZone(zone int) (Fragment, bool) {
	for _, f := range Fragments {
		if f.Zone == zone {
			return f, true
		}
	}
	return Fragment{}, false
}
