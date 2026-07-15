package catalog

import (
	"fmt"

	"game/internal/entity"
)

// Балансные параметры бонусов
const (
	rageDamageBonus = 2 // насколько ярость поднимает урон, пока не спала
	rageDuration    = 3 // сколько ходов держится ярость
	ragePenalty     = 1 // насколько ярость режет урон после того, как спала

	shieldReduction   = 3 // насколько щит режет входящий урон
	hiddenAttackBonus = 1 // прибавка скрытой атаки при перевесе в ловкости
	impulseMultiplier = 2 // во сколько раз порыв к действию множит урон в первый ход
	poisonPerRound    = 1 // насколько яд наращивает урон с каждым ходом

	attributeGainStep = 1 // на сколько бонусы вида "+1 к атрибуту" двигают атрибут
)

// Bonuses - таблица бонусов (идентификатор -> что бонус собой представляет и что делает)
var Bonuses = map[entity.BonusName]entity.Bonus{
	entity.Rage: {
		Title:       "Ярость",
		Description: fmt.Sprintf("+%d к урону в первые %d хода, потом −%d к урону", rageDamageBonus, rageDuration, ragePenalty),
		Outcome: func(ctx entity.FightContext, damage uint16) uint16 {
			if ctx.Round <= rageDuration {
				return damage + rageDamageBonus
			}
			return damageChecker(damage, ragePenalty)
		},
	},
	entity.StoneSkin: {
		Title:       "Каменная кожа",
		Description: "получаемый урон снижается на значение выносливости",
		Income: func(ctx entity.FightContext, damage uint16) uint16 {
			return damageChecker(damage, uint16(ctx.Defender.Attributes.Stamina))
		},
	},
	entity.HiddenAttack: {
		Title:       "Скрытая атака",
		Description: fmt.Sprintf("+%d к урону, если ловкость персонажа выше ловкости цели", hiddenAttackBonus),
		Outcome: func(ctx entity.FightContext, damage uint16) uint16 {
			if ctx.Attacker.Attributes.Agility > ctx.Defender.Attributes.Agility {
				return damage + hiddenAttackBonus
			}
			return damage
		},
	},
	entity.Shield: {
		Title:       "Щит",
		Description: fmt.Sprintf("−%d к получаемому урону, если сила персонажа выше силы атакующего", shieldReduction),
		Income: func(ctx entity.FightContext, damage uint16) uint16 {
			if ctx.Defender.Attributes.Strength > ctx.Attacker.Attributes.Strength {
				return damageChecker(damage, shieldReduction)
			}
			return damage
		},
	},
	entity.ImpulseToAction: {
		Title:       "Порыв к действию",
		Description: fmt.Sprintf("в первый ход урон оружием умножается на %d", impulseMultiplier),
		Outcome: func(ctx entity.FightContext, damage uint16) uint16 {
			if ctx.Round == 1 {
				return damage * impulseMultiplier
			}
			return damage
		},
	},
	entity.Poison: {
		Title: "Яд",
		Description: fmt.Sprintf("доп. +%d урона на 2-м ходу, +%d на 3-м и так далее",
			poisonPerRound, poisonPerRound*2),
		Outcome: func(ctx entity.FightContext, damage uint16) uint16 {
			if ctx.Round >= 2 {
				return damage + uint16(ctx.Round-1)*poisonPerRound
			}
			return damage
		},
	},
	entity.AgilityUp: {
		Title:       fmt.Sprintf("Ловкость +%d", attributeGainStep),
		Description: fmt.Sprintf("+%d к ловкости", attributeGainStep),
		AttributeGain: func(h *entity.Hero) {
			h.Attributes.Agility += attributeGainStep
		},
	},
	entity.StrengthUp: {
		Title:       fmt.Sprintf("Сила +%d", attributeGainStep),
		Description: fmt.Sprintf("+%d к силе", attributeGainStep),
		AttributeGain: func(h *entity.Hero) {
			h.Attributes.Strength += attributeGainStep
		},
	},
	entity.StaminaUp: {
		Title:       fmt.Sprintf("Выносливость +%d", attributeGainStep),
		Description: fmt.Sprintf("+%d к выносливости", attributeGainStep),
		AttributeGain: func(h *entity.Hero) {
			h.Attributes.Stamina += attributeGainStep
		},
	},
}

// damageChecker вычитает resist из урона с полом в 0 (беззнаковый тип не уходит в минус).
func damageChecker(dmg, resist uint16) uint16 {
	if resist >= dmg {
		return 0
	}
	return dmg - resist
}
