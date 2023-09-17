package handlers

import (
	"database/sql"
	"net/http"
	"sarkaz_api/database"
	"sarkaz_api/models"

	"github.com/gin-gonic/gin"
)

func getCharacterByName(db *sql.DB, name string) (models.Character, error) {
    var character models.Character
    err := db.QueryRow("SELECT * FROM Characters WHERE name=?", name).Scan(
        &character.CharacterID, &character.Name, &character.Level, &character.EliteLevel,
        &character.MaxHP, &character.Attack, &character.Defense, &character.MagicResistance,
        &character.RedeployTime, &character.DeployCost, &character.BlockCount,
        &character.AttackInterval, &character.Rarity,
    )
    if err != nil {
        return character, err
    }
    return character, nil
}

func getSimpleCharacterByName(db *sql.DB, name string) (models.CharacterSimple, error) {
    var character models.CharacterSimple
    err := db.QueryRow("SELECT Name, Rarity FROM Characters WHERE name=?", name).Scan(
         &character.Name, &character.Rarity,
    )
    if err != nil {
        return character, err
    }
    return character, nil
}


func GetCharacterHandler(c *gin.Context) {
    characterName := c.Param("name")
    character, err := getCharacterByName(database.DB, characterName)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
        return
    }
    c.JSON(http.StatusOK, character)
}

func GetSimpleCharacterHandler(c *gin.Context) {
    characterName := c.Param("name")
    character, err := getSimpleCharacterByName(database.DB, characterName)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
        return
    }
    c.JSON(http.StatusOK, character)
}