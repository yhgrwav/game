package catalog

import "game/internal/entity"

// Enemies - таблица монстров
var Enemies = []entity.Enemy{
	{
		Name:   "Гоблин",
		HP:     5,
		Damage: 2,
		Stats: entity.Attributes{
			Strength: 1,
			Agility:  1,
			Stamina:  1,
		},
		Feature: entity.FeatureNone,
		Reward:  Dagger,
	},
	{
		Name:   "Скелет",
		HP:     10,
		Damage: 2,
		Stats: entity.Attributes{
			Strength: 2,
			Agility:  2,
			Stamina:  1,
		},
		Feature: entity.SkeletonWeakCrushing,
		Reward:  Club,
	},
	{
		Name:   "Слайм",
		HP:     8,
		Damage: 1,
		Stats: entity.Attributes{
			Strength: 3,
			Agility:  1,
			Stamina:  2,
		},
		Feature: entity.SlimeImmuneChopping,
		Reward:  Spear,
	},
	{
		Name:   "Призрак",
		HP:     6,
		Damage: 3,
		Stats: entity.Attributes{
			Strength: 1,
			Agility:  3,
			Stamina:  1,
		},
		Feature: entity.GhostHiddenAttack,
		Reward:  Sword,
	},
	{
		Name:   "Голем",
		HP:     10,
		Damage: 1,
		Stats: entity.Attributes{
			Strength: 3,
			Agility:  1,
			Stamina:  3,
		},
		Feature: entity.GolemStoneSkin,
		Reward:  Axe,
	},
	{
		Name:   "Дракон",
		HP:     20,
		Damage: 4,
		Stats: entity.Attributes{
			Strength: 3,
			Agility:  3,
			Stamina:  3,
		},
		Feature: entity.DragonBreath,
		Reward:  LegendarySword,
	},
}
