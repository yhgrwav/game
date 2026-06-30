package entity

type damageType int

const (
	unknown damageType = iota

	// Рубящий
	chopping

	// Дробящий
	crushing

	// Колющий
	pricking
)

type Weapon struct {
	Name    string
	Damage  uint16
	DmgType damageType
}

// Таблица оружия.
var (
	Sword = Weapon{
		Name:    "Sword",
		Damage:  3,
		DmgType: chopping,
	}
	Club = Weapon{
		Name:    "club",
		Damage:  3,
		DmgType: crushing,
	}
	Dagger = Weapon{
		Name:    "dagger",
		Damage:  2,
		DmgType: pricking,
	}
	Axe = Weapon{
		Name:    "axe",
		Damage:  4,
		DmgType: chopping,
	}
	Spear = Weapon{
		Name:    "spear",
		Damage:  3,
		DmgType: pricking,
	}
	LegendarySword = Weapon{
		Name:    "legendary sword",
		Damage:  10,
		DmgType: chopping,
	}
)
