package events

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"game/internal/catalog"
	"game/internal/entity"
	"game/internal/methods"
)

// AfterVictory: хил до максимума, предложение прокачки и замены оружия
func AfterVictory(hero *entity.Hero, enemy entity.Enemy) {
	hero.HP = hero.MaxHP
	fmt.Printf("HP восстановлено: %d/%d\n", hero.HP, hero.MaxHP)
	offerLevelUp(hero)
	offerWeapon(hero, enemy.Reward)
}

func offerLevelUp(hero *entity.Hero) {
	if methods.TotalLevel(hero) >= 3 {
		fmt.Println("Максимальный уровень (3) достигнут.")
		return
	}
	fmt.Println("Повысить уровень? Выберите класс (0 — пропустить):")
	for i, c := range catalog.Classes {
		fmt.Printf("%d.%s\n", i+1, c.Title)
	}
	idx, err := strconv.Atoi(readLine())
	if err != nil || idx < 1 || idx > len(catalog.Classes) {
		return
	}
	methods.LevelUp(hero, catalog.Classes[idx-1].Class)
	fmt.Printf("Теперь: %s (%d/%d HP)\n", methods.HeroTitle(hero), hero.HP, hero.MaxHP)
}

func offerWeapon(hero *entity.Hero, drop entity.Weapon) {
	fmt.Printf("Выпало оружие: %s (%d урона). Заменить? (1 — да, иначе нет)\n", drop.Name, drop.Damage)
	if readLine() == "1" {
		hero.Weapon = drop
		fmt.Printf("Оружие заменено на %s.\n", drop.Name)
	}
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}
