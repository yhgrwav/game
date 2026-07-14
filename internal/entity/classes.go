package entity

type HeroClass string

const (
	// UnknownHeroClass - заглушка для невалидного ввода
	UnknownHeroClass HeroClass = ""

	// Robber - разбойник.
	Robber HeroClass = "Robber"

	// Warrior - воин.
	Warrior HeroClass = "Warrior"

	// Barbarian - варвар.
	Barbarian HeroClass = "Barbarian"
)

type ClassInfo struct {
	Class HeroClass
	Title string
}

// HeroParams - шаблон параметров героя
type HeroParams struct {
	BasicAttributes Attributes
	HpPerLevel      uint8
	Weapon          Weapon
	BonusOne        BonusName
	BonusTwo        BonusName
	BonusThree      BonusName
}
