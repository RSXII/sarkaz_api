package models

type IDNumber int

type UnitBaseStats struct {
    Name             string
    Level            int
    MaxHP            int
    Attack           int
    Defense          int
    MagicResistance  int
    AttackInterval   float64
}
type Character struct {
    UnitBaseStats
    CharacterID      IDNumber
    RedeployTime     int
    DeployCost       int
    BlockCount       int
    Rarity           int
    EliteLevel       int
}

type Enemy struct {
    UnitBaseStats
    EnemyId         IDNumber
}

type CharacterSimple struct {
    Name             string
    Rarity           int
}

type TrustBonus struct {
    TrustBonusID     int
    CharacterID      int
    StatToIncrease   string
    Amount           int
}

type CharacterCollection struct {
    CollectionId         int
    CollectionList       []Character
}

type Unit struct {
    unitType    string
}