package database

import (
	"database/sql"
	"sarkaz_api/models"
)

type CharacterRepository struct {
    db *sql.DB
}

func NewCharacterRepository(db *sql.DB) *CharacterRepository {
    return &CharacterRepository{db}
}

func (repo *CharacterRepository) GetCharacterByName(name string) (models.Character, error) {
    var character models.Character
    err := repo.db.QueryRow("SELECT * FROM Characters WHERE name=?", name).Scan(
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

func (repo *CharacterRepository) GetSimpleCharacterByName(name string) (models.CharacterSimple, error) {
    var character models.CharacterSimple
    err := repo.db.QueryRow("SELECT Name, Rarity FROM Characters WHERE name=?", name).Scan(
        &character.Name, &character.Rarity,
    )
    if err != nil {
        return character, err
    }
    return character, nil
}
