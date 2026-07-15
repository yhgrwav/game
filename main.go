package main

import (
	"fmt"

	"game/internal/events"
)

const winsToFinish = 5

func main() {
	for {
		hero := events.StartGame()
		if hero == nil {
			return
		}

		wins := 0
		for {
			enemy, won := events.RunFight(hero)
			if !won {
				fmt.Println("Поражение. Создаём нового персонажа.")
				break
			}

			wins++
			fmt.Printf("Побед подряд: %d/%d\n", wins, winsToFinish)
			if wins == winsToFinish {
				fmt.Println("Игра пройдена! 🏆")
				return
			}

			events.AfterVictory(hero, enemy)
		}
	}
}
