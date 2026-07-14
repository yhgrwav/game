package catalog

import "game/internal/entity"

// Таблица оружия
var (
	Sword = entity.Weapon{
		Name:    "Меч",
		Damage:  3,
		DmgType: entity.Chopping,
	}
	Club = entity.Weapon{
		Name:    "Дубина",
		Damage:  3,
		DmgType: entity.Crushing,
	}
	Dagger = entity.Weapon{
		Name:    "Кинжал",
		Damage:  2,
		DmgType: entity.Pricking,
	}
	Axe = entity.Weapon{
		Name:    "Топор",
		Damage:  4,
		DmgType: entity.Chopping,
	}
	Spear = entity.Weapon{
		Name:    "Копьё",
		Damage:  3,
		DmgType: entity.Pricking,
	}
	LegendarySword = entity.Weapon{
		Name:    "Легендарный меч",
		Damage:  10,
		DmgType: entity.Chopping,
	}
)
