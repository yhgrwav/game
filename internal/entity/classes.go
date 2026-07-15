package entity

type HeroClass string

const (
	Robber    HeroClass = "Robber"
	Warrior   HeroClass = "Warrior"
	Barbarian HeroClass = "Barbarian"
)

type ClassInfo struct {
	Class HeroClass
	Title string
}

// ClassTemplate - статика класса без рандома, одинаковая для всех героев класса.
// Bonuses упорядочены по уровню: индекс = уровень-1.
type ClassTemplate struct {
	HpPerLevel uint8
	Weapon     Weapon
	Bonuses    [3]BonusName
}
