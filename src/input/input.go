// Package input centralise la lecture des entrées clavier.
// Un seul lecteur partagé évite les conflits de buffer entre packages,
// et permettra plus tard de rediriger l'entrée (tests, future TUI).
package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// Line affiche un prompt optionnel et renvoie la ligne saisie, nettoyée.
func Line(prompt string) string {
	if prompt != "" {
		fmt.Print(prompt)
	}
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// Int lit un entier. Renvoie (0, false) si la saisie n'est pas un entier valide.
func Int(prompt string) (int, bool) {
	n, err := strconv.Atoi(Line(prompt))
	if err != nil {
		return 0, false
	}
	return n, true
}
