package entity

// Hero - персонаж игрока: ClassLevels хранит уровень по каждому классу
// (мультикласс), суммарный уровень - сумма значений, максимум 3
type Hero struct {
	Attributes  Attributes
	Weapon      Weapon
	Bonuses     []BonusName
	ClassLevels map[HeroClass]uint8
	HP          uint8
	MaxHP       uint8
}
