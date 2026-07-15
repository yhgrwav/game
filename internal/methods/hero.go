package methods

import (
	"game/internal/catalog"
	"game/internal/entity"
)

func NewHero(class entity.HeroClass) *entity.Hero {
	build, ok := catalog.Heroes[class]
	if !ok {
		return nil // невалидный класс
	}
	params := build()
	maxHP := params.HpPerLevel + uint8(params.BasicAttributes.Stamina)

	return &entity.Hero{
		Attributes: params.BasicAttributes,
		Class:      class,
		Weapon:     params.Weapon,
		Bonuses: []entity.BonusName{
			params.BonusOne,
			params.BonusTwo,
			params.BonusThree,
		},
		Level:      1,
		HP:         maxHP,
		MaxHP:      maxHP,
		HpPerLevel: params.HpPerLevel,
	}
}
