package main

import (
	"bufio"
	"fmt"
	"os"

	"eldoria/interne/game"
)

func main() {
	fmt.Println("================================")
	fmt.Println("ELDORIA")
	fmt.Println("Les Fragments du Souvenir")
	fmt.Println("================================")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	state := game.NewGame(scanner)

	game.Run(scanner, &state)
}
