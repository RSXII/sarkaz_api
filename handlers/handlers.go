package handlers

import (
	"net/http"
	"sarkaz_api/auth"
	"sarkaz_api/database"
	"sarkaz_api/models"

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

func SignUpHandler(usersRepo *database.UserRepository) func (c *gin.Context) {
    return func(c *gin.Context) {
        var userData models.User
        err := c.BindJSON(userData)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user data"})
            return
        }

        token, err := auth.CreateToken(userData.UserId)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"token": token})
    }
}