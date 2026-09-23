package player

import "testing"

// Cas du sujet : besoin 100, XP 90, gain +30 -> niveau +1, reste 20.
func TestGainXPKeepsOverflow(t *testing.T) {
	c := InitCharacter()
	c.MaxXP = 100
	c.XP = 90

	GainXP(&c, 30)

	if c.Level != 2 {
		t.Fatalf("niveau attendu 2, obtenu %d", c.Level)
	}
	if c.XP != 20 {
		t.Fatalf("XP restante attendue 20, obtenue %d", c.XP)
	}
}

// Un gain massif doit pouvoir enchaîner plusieurs niveaux.
func TestGainXPMultipleLevels(t *testing.T) {
	c := InitCharacter()
	c.MaxXP = 100
	c.XP = 0

	// 100 (lvl2, seuil->150) + 150 (lvl3, seuil->200) + 50 restants.
	GainXP(&c, 300)

	if c.Level != 3 {
		t.Fatalf("niveau attendu 3, obtenu %d", c.Level)
	}
	if c.XP != 50 {
		t.Fatalf("XP restante attendue 50, obtenue %d", c.XP)
	}
}

// L'XP négative ou nulle ne doit rien changer.
func TestGainXPIgnoresNonPositive(t *testing.T) {
	c := InitCharacter()
	c.MaxXP = 100
	c.XP = 40

	GainXP(&c, -50)
	GainXP(&c, 0)

	if c.XP != 40 || c.Level != 1 {
		t.Fatalf("état inattendu : XP=%d Level=%d", c.XP, c.Level)
	}
}

// Revive ne ranime qu'un personnage mort, à 50 % des PV max.
func TestRevive(t *testing.T) {
	c := InitCharacter()
	c.MaxHP = 100
	c.HP = 0

	if !Revive(&c) {
		t.Fatal("Revive aurait dû ranimer un personnage mort")
	}
	if c.HP != 50 {
		t.Fatalf("PV attendus 50, obtenus %d", c.HP)
	}
	if Revive(&c) {
		t.Fatal("Revive ne devrait rien faire sur un personnage vivant")
	}
}

// Heal ne dépasse jamais les PV max.
func TestHealCapsAtMax(t *testing.T) {
	c := InitCharacter()
	c.MaxHP = 100
	c.HP = 80

	Heal(&c, 50)

	if c.HP != 100 {
		t.Fatalf("PV attendus 100 (plafonné), obtenus %d", c.HP)
	}
}
