package methods

import (
	"game/internal/catalog"
	"game/internal/entity"
)

// applyDamage применяет урон к хп дефендера: возвращает остаток хп (с полом в 0, чтобы беззнаковый uint16
// не ушёл в минус и не превратился в несколько тысяч) и флаг смерти
func applyDamage(dmg, health uint16) (remaining uint16, isDead bool) {
	// снесли в ноль или ниже - дефендер мёртв, остаток хп = 0
	if dmg >= health {
		return 0, true
	}
	// урон пережит - возвращаем остаток хп
	return health - dmg, false
}

// Fight считает итоговый урон удара attacker -> defender на ходу round
func Fight(attacker, defender *entity.Hero, round uint8) uint16 {
	ctx := entity.FightContext{
		Attacker: attacker,
		Defender: defender,
		Round:    round,
	}

	dmg := attacker.Weapon.Damage

	for _, name := range attacker.Bonuses {
		if b := catalog.Bonuses[name]; b.Outcome != nil {
			dmg = b.Outcome(ctx, dmg)
		}
	}

	for _, name := range defender.Bonuses {
		if b := catalog.Bonuses[name]; b.Income != nil {
			dmg = b.Income(ctx, dmg)
		}
	}

	return dmg
}
