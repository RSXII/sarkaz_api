package character_service

import (
	"fmt"
	"sarkaz_api/models"
)


func OutputCharactersFromList(characterList []models.Character) {
	for _, character := range characterList {
		fmt.Println(character.Name)
	}
}