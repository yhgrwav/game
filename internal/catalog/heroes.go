package catalog

import (
	"math/rand"

	"game/internal/entity"
)

// Templates - статика классов (без атрибутов, они случайны на каждого героя)
var Templates = map[entity.HeroClass]entity.ClassTemplate{
	entity.Robber: {
		HpPerLevel: 4,
		Weapon:     Dagger,
		Bonuses:    [3]entity.BonusName{entity.HiddenAttack, entity.AgilityUp, entity.Poison},
	},
	entity.Warrior: {
		HpPerLevel: 5,
		Weapon:     Sword,
		Bonuses:    [3]entity.BonusName{entity.ImpulseToAction, entity.Shield, entity.StrengthUp},
	},
	entity.Barbarian: {
		HpPerLevel: 6,
		Weapon:     Club,
		Bonuses:    [3]entity.BonusName{entity.Rage, entity.StoneSkin, entity.StaminaUp},
	},
}

// RollAttributes перекатывает атрибуты нового героя (каждый 1-3)
func RollAttributes() entity.Attributes {
	return entity.Attributes{
		Strength: entity.Strength(randomAttribute()),
		Agility:  entity.Agility(randomAttribute()),
		Stamina:  entity.Stamina(randomAttribute()),
	}
}

func randomAttribute() uint16 {
	return uint16(rand.Intn(3) + 1)
}
