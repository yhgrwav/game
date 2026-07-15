package methods

import (
	"math/rand"

	"game/internal/catalog"
	"game/internal/entity"
)

const (
	dragonFireBonus = 3 // доп. урон дыхания дракона
	dragonFireEvery = 3 // дракон дышит огнём каждый N-й ход
)

// HeroStrikesFirst - первым бьёт тот, у кого выше ловкость; при равенстве герой
func HeroStrikesFirst(hero, enemy *entity.Combatant) bool {
	return hero.Attributes.Agility >= enemy.Attributes.Agility
}

// Hits - бросок на попадание
func Hits(attacker, defender *entity.Combatant) bool {
	atk := int(attacker.Attributes.Agility)
	def := int(defender.Attributes.Agility)
	roll := rand.Intn(atk+def) + 1
	return roll > def
}

// Attack проводит удар: проверяет попадание, считает урон и вычитает его из HP
// цели. Возвращает нанесённый урон и флаг попадания
func Attack(attacker, defender *entity.Combatant, round uint8) (damage uint16, hit bool) {
	if !Hits(attacker, defender) {
		return 0, false
	}
	damage = strike(attacker, defender, round)
	defender.HP = applyDamage(damage, defender.HP)
	return damage, true
}

// strike считает итоговый урон удара без учёта попадания и без изменения HP
func strike(attacker, defender *entity.Combatant, round uint8) uint16 {
	ctx := entity.FightContext{Attacker: attacker, Defender: defender, Round: round}

	weaponDmg := attacker.Weapon.Damage
	switch defender.Feature {
	case entity.SkeletonWeakCrushing:
		if attacker.Weapon.DmgType == entity.Crushing {
			weaponDmg *= 2
		}
	case entity.SlimeImmuneChopping:
		if attacker.Weapon.DmgType == entity.Chopping {
			weaponDmg = 0
		}
	}

	dmg := weaponDmg + uint16(attacker.Attributes.Strength)

	for _, name := range attacker.Bonuses {
		if b := catalog.Bonuses[name]; b.Outcome != nil {
			dmg = b.Outcome(ctx, dmg)
		}
	}

	if attacker.Feature == entity.DragonBreath && round%dragonFireEvery == 0 {
		dmg += dragonFireBonus
	}

	for _, name := range defender.Bonuses {
		if b := catalog.Bonuses[name]; b.Income != nil {
			dmg = b.Income(ctx, dmg)
		}
	}

	return dmg
}

// applyDamage вычитает урон из HP с полом в 0, чтобы беззнаковый тип не переполнился
func applyDamage(dmg, hp uint16) uint16 {
	if dmg >= hp {
		return 0
	}
	return hp - dmg
}
