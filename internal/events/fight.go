package events

import (
	"fmt"
	"math/rand"
	"time"

	"game/internal/catalog"
	"game/internal/entity"
	"game/internal/methods"
)

// strikeDelay - пауза между ударами
const strikeDelay = 900 * time.Millisecond

// RunFight проводит пошаговый бой героя со случайным монстром
func RunFight(hero *entity.Hero) (entity.Enemy, bool) {
	enemy := catalog.Enemies[rand.Intn(len(catalog.Enemies))]
	hc := methods.NewCombatantFromHero(hero)
	ec := methods.NewCombatantFromEnemy(&enemy)

	fmt.Printf("\n%s (%d HP) против %s (%d HP)\n", hc.Name, hc.HP, ec.Name, ec.HP)
	time.Sleep(strikeDelay)

	var heroRound, enemyRound uint8
	heroFirst := methods.HeroStrikesFirst(hc, ec)

	for hc.HP > 0 && ec.HP > 0 {
		if heroFirst {
			heroRound++
			logStrike(hc, ec, heroRound)
			if ec.HP == 0 {
				break
			}
			enemyRound++
			logStrike(ec, hc, enemyRound)
		} else {
			enemyRound++
			logStrike(ec, hc, enemyRound)
			if hc.HP == 0 {
				break
			}
			heroRound++
			logStrike(hc, ec, heroRound)
		}
	}

	heroWon := hc.HP > 0
	if heroWon {
		fmt.Printf("%s побеждает!\n", hc.Name)
	} else {
		fmt.Printf("%s погибает.\n", hc.Name)
	}
	return enemy, heroWon
}

func logStrike(attacker, defender *entity.Combatant, round uint8) {
	dmg, hit := methods.Attack(attacker, defender, round)
	if !hit {
		fmt.Printf("  %s промахивается по %s\n", attacker.Name, defender.Name)
	} else {
		fmt.Printf("  %s бьёт %s на %d (у %s осталось %d HP)\n",
			attacker.Name, defender.Name, dmg, defender.Name, defender.HP)
	}
	time.Sleep(strikeDelay)
}
