package main

import (
	"eldoria/entity"
)

func main() {
	perso := entity.CharacterCreation()
	entity.DisplayInfo(perso)
}
