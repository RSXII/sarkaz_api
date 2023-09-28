package handlers

import (
	"net/http"
	"sarkaz_api/database"

	"github.com/gin-gonic/gin"
)

func GetCharacterHandler(characterRepo *database.CharacterRepository) func(c *gin.Context) {
    return func(c *gin.Context) {
        characterName := c.Param("name")
        character, err := characterRepo.GetCharacterByName(characterName)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
            return
        }
        c.JSON(http.StatusOK, character)
    }
}

func GetSimpleCharacterHandler(characterRepo *database.CharacterRepository) func(c *gin.Context) {
    return func(c *gin.Context) {
        characterName := c.Param("name")
        character, err := characterRepo.GetSimpleCharacterByName(characterName)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
        return
    }
    c.JSON(http.StatusOK, character)
    }
}

func GetCharactersByRarityHandler(characterRepo *database.CharacterRepository) func(c *gin.Context) {
    return func(c *gin.Context){
        characterRating := c.Param("stars")
        characters, err := characterRepo.GetCharactersByRarity(characterRating)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
    }
    c.JSON(http.StatusOK, characters)
    }
    
}