package main

import (
	"eldoria/entity"
	"fmt"
)

func main() {
	perso := entity.CharacterCreation()
	fmt.Println(perso.Name, perso.Class, perso.HP)
}
