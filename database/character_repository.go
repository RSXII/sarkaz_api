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
    query := "SELECT * FROM Characters WHERE name=?"

    stmt, error := repo.db.Prepare(query)
    if error != nil {
        return character, error
    }
    defer stmt.Close()
    err := stmt.QueryRow(name).Scan(
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
    query := "SELECT Name, Rarity FROM Characters WHERE name=?"

    stmt, error := repo.db.Prepare(query)
    if error != nil {
        return character, nil
    }
    defer stmt.Close()
    err := stmt.QueryRow(name).Scan(
        &character.Name, &character.Rarity,
    )
    if err != nil {
        return character, err
    }
    return character, nil
}

func (repo *CharacterRepository) GetCharactersByRarity(rarity string) ([]models.Character, error) {
	var characters []models.Character
    query := "SELECT * FROM Characters WHERE rarity=?"

    stmt, error := repo.db.Prepare(query)
    if error != nil {
        return nil, error
    }
    defer stmt.Close()

    rows, error := stmt.Query(rarity)
    if error != nil {
        return nil, error
    }

    for rows.Next() {
        var character models.Character

        err := rows.Scan(
            &character.CharacterID, &character.Name, &character.Level, &character.EliteLevel,
            &character.MaxHP, &character.Attack, &character.Defense, &character.MagicResistance,
            &character.RedeployTime, &character.DeployCost, &character.BlockCount,
            &character.AttackInterval, &character.Rarity,
        )
        if err != nil {
            return nil, err
        }
        characters = append(characters, character)
    }

	if err := rows.Err(); err != nil {
        return nil, err
    }
	return characters, nil
}
