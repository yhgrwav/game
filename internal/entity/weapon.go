package entity

type DamageType int

const (
	// UnknownDamage - нулевое значение, осмысленного оружия с таким типом не бывает.
	UnknownDamage DamageType = iota

	// Chopping - рубящий.
	Chopping

	// Crushing - дробящий.
	Crushing

	// Pricking - колющий.
	Pricking
)

type Weapon struct {
	Name    string
	Damage  uint16
	DmgType DamageType
}
