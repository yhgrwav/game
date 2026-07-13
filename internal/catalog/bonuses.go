package catalog

import "game/internal/entity"

// Bonuses - таблица бонусов.
//
// Обычная (не ленивая) мапа: бонусы статичны и не хранят состояния конкретного героя,
// всё нужное Income/Outcome берут из FightContext в момент вызова. Пересоздавать их на каждое
// обращение незачем, поэтому функция-конструктор в значении здесь (в отличие от Heroes) не нужна.
// Их читают на каждый удар в бою, так что лишние аллокации тут были бы особенно некстати.
var Bonuses = map[entity.BonusName]entity.Bonus{
	entity.Rage: {
		Description: "+2 к урону в первые 3 хода, потом −1 к урону",
		Outcome: func(ctx entity.FightContext, damage uint16) uint16 {
			if ctx.Round <= 3 {
				return damage + 2
			}
			return damageChecker(damage, 1)
		},
	},
	entity.StoneSkin: {
		Description: "получаемый урон снижается на значение выносливости",
		Income: func(ctx entity.FightContext, damage uint16) uint16 {
			return damageChecker(damage, uint16(ctx.Defender.Attributes.Stamina))
		},
	},
	entity.HiddenAttack: {
		Description: "+1 к урону, если ловкость персонажа выше ловкости цели",
		Outcome: func(ctx entity.FightContext, damage uint16) uint16 {
			if ctx.Attacker.Attributes.Agility > ctx.Defender.Attributes.Agility {
				return damage + 1
			}
			return damage
		},
	},
	entity.Shield: {
		Description: "−3 к получаемому урону, если сила персонажа выше силы атакующего",
		Income: func(ctx entity.FightContext, damage uint16) uint16 {
			if ctx.Defender.Attributes.Strength > ctx.Attacker.Attributes.Strength {
				return damageChecker(damage, 3)
			}
			return damage
		},
	},
	entity.ImpulseToAction: {
		Description: "в первый ход двойной урон оружием",
		Outcome: func(ctx entity.FightContext, damage uint16) uint16 {
			if ctx.Round == 1 {
				return damage * 2
			}
			return damage
		},
	},
	entity.Poison: {
		Description: "доп. +1 урона на 2-м ходу, +2 на 3-м и т.д.",
		Outcome: func(ctx entity.FightContext, damage uint16) uint16 {
			if ctx.Round >= 2 {
				return damage + uint16(ctx.Round-1)
			}
			return damage
		},
	},
	entity.AgilityUp: {
		Description: "+1 к ловкости",
		AttributeGain: func(h *entity.Hero) {
			h.Attributes.Agility++
		},
	},
	entity.StrengthUp: {
		Description: "+1 к силе",
		AttributeGain: func(h *entity.Hero) {
			h.Attributes.Strength++
		},
	},
	entity.StaminaUp: {
		Description: "+1 к выносливости",
		AttributeGain: func(h *entity.Hero) {
			h.Attributes.Stamina++
		},
	},
}

// damageChecker безопасно вычитает resist из урона: при resist >= dmg возвращает 0,
// чтобы беззнаковый uint16 не ушёл в минус и не превратился в десятки тысяч.
// Живёт рядом с Bonuses, потому что нужен только их замыканиям: унеси его в methods -
// и получишь цикл импортов (methods уже импортирует catalog).
func damageChecker(dmg, resist uint16) uint16 {
	if resist >= dmg {
		return 0
	}
	return dmg - resist
}
