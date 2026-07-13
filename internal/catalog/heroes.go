package catalog

import (
	"math/rand"

	"game/internal/entity"
)

// Heroes - таблица классов
//
// ключ - HeroClass, значение - конструктор для класса
var Heroes = map[entity.HeroClass]func() entity.HeroParams{
	entity.Robber: func() entity.HeroParams {
		return entity.HeroParams{
			BasicAttributes: rollAttributes(),
			HpPerLevel:      4,
			Weapon:          Dagger,
			BonusOne:        entity.HiddenAttack,
			BonusTwo:        entity.AgilityUp,
			BonusThree:      entity.Poison,
		}
	},
	entity.Warrior: func() entity.HeroParams {
		return entity.HeroParams{
			BasicAttributes: rollAttributes(),
			HpPerLevel:      5,
			Weapon:          Sword,
			BonusOne:        entity.ImpulseToAction,
			BonusTwo:        entity.Shield,
			BonusThree:      entity.StrengthUp,
		}
	},
	entity.Barbarian: func() entity.HeroParams {
		return entity.HeroParams{
			BasicAttributes: rollAttributes(),
			HpPerLevel:      6,
			Weapon:          Club,
			BonusOne:        entity.Rage,
			BonusTwo:        entity.StoneSkin,
			BonusThree:      entity.StaminaUp,
		}
	},
}

// rollAttributes перекатывает все три атрибута нового героя
func rollAttributes() entity.Attributes {
	return entity.Attributes{
		Strength: entity.Strength(randomAttribute()),
		Agility:  entity.Agility(randomAttribute()),
		Stamina:  entity.Stamina(randomAttribute()),
	}
}

// randomAttribute возвращает случайное число от 1 до 3 включительно
func randomAttribute() uint16 {
	return uint16(rand.Intn(3) + 1)
}
