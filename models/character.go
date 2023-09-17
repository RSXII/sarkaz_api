package models

type Character struct {
    CharacterID      int
    Name             string
    Level            int
    EliteLevel       int
    MaxHP            int
    Attack           int
    Defense          int
    MagicResistance  int
    RedeployTime     int
    DeployCost       int
    BlockCount       int
    AttackInterval   float64
    Rarity           int
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
