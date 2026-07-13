package entity

type Hero struct {
	Attributes Attributes
	Class      HeroClass
	Weapon     Weapon
	Bonuses    []BonusName
	Level      uint8
	HP         uint8
	MaxHP      uint8
	HpPerLevel uint8
}
