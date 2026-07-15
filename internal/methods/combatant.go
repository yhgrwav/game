package methods

import "game/internal/entity"

func NewCombatantFromHero(h *entity.Hero) *entity.Combatant {
	return &entity.Combatant{
		Name:       HeroTitle(h),
		Attributes: h.Attributes,
		Weapon:     h.Weapon,
		Bonuses:    h.Bonuses,
		Feature:    entity.FeatureNone,
		HP:         uint16(h.HP),
	}
}

func NewCombatantFromEnemy(e *entity.Enemy) *entity.Combatant {
	return &entity.Combatant{
		Name:       e.Name,
		Attributes: e.Stats,
		Weapon:     entity.Weapon{Damage: uint16(e.Damage)},
		Bonuses:    enemyBonuses(e.Feature),
		Feature:    e.Feature,
		HP:         uint16(e.HP),
	}
}

// enemyBonuses раскрывает особенности монстра, совпадающие с геройскими бонусами,
// в список бонусов; остальные разбирает боевая логика по Feature
func enemyBonuses(f entity.FeatureKind) []entity.BonusName {
	switch f {
	case entity.GhostHiddenAttack:
		return []entity.BonusName{entity.HiddenAttack}
	case entity.GolemStoneSkin:
		return []entity.BonusName{entity.StoneSkin}
	default:
		return nil
	}
}
