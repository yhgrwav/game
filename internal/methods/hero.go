package methods

import "game/internal/entity"

func NewHero(class entity.HeroClass) *entity.Hero {
	build, ok := entity.Heroes[class]
	if !ok {
		return nil // невалидный класс
	}
	params := build()
	return &entity.Hero{
		Attributes: params.BasicAttributes,
		Class:      class,
		Weapon:     params.Weapon,
		// бонусы прокидываем в самого героя - именно отсюда Fight их потом перебирает
		Bonuses: []entity.Bonus{
			params.BonusOne,
			params.BonusTwo,
			params.BonusThree,
		},
	}
}
