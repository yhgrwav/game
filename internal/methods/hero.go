package methods

import (
	"fmt"
	"strings"

	"game/internal/catalog"
	"game/internal/entity"
)

// NewHero создаёт героя 1 уровня выбранного класса: катает атрибуты и повышает
// уровень стартовым классом, чтобы создание и прокачка шли одним путём
func NewHero(class entity.HeroClass) *entity.Hero {
	if _, ok := catalog.Templates[class]; !ok {
		return nil
	}
	h := &entity.Hero{
		Attributes:  catalog.RollAttributes(),
		ClassLevels: map[entity.HeroClass]uint8{},
	}
	LevelUp(h, class)
	return h
}

// LevelUp повышает уровень героя в классе: +1 уровень -> бонус этого уровня ->
// возможная прибавка атрибута -> рост MaxHP -> хил до полного
func LevelUp(h *entity.Hero, class entity.HeroClass) {
	tmpl, ok := catalog.Templates[class]
	if !ok {
		return
	}
	h.ClassLevels[class]++
	level := h.ClassLevels[class]

	if int(level) <= len(tmpl.Bonuses) {
		name := tmpl.Bonuses[level-1]
		h.Bonuses = append(h.Bonuses, name)
		if b := catalog.Bonuses[name]; b.AttributeGain != nil {
			b.AttributeGain(h)
		}
	}

	if h.Weapon == (entity.Weapon{}) {
		h.Weapon = tmpl.Weapon
	}

	h.MaxHP += tmpl.HpPerLevel + uint8(h.Attributes.Stamina)
	h.HP = h.MaxHP
}

// TotalLevel - суммарный уровень героя по всем классам
func TotalLevel(h *entity.Hero) uint8 {
	var total uint8
	for _, lvl := range h.ClassLevels {
		total += lvl
	}
	return total
}

// HeroTitle - читаемое имя героя вида "Воин 1 / Разбойник 1"
func HeroTitle(h *entity.Hero) string {
	var parts []string
	for _, c := range catalog.Classes {
		if lvl := h.ClassLevels[c.Class]; lvl > 0 {
			parts = append(parts, fmt.Sprintf("%s %d", c.Title, lvl))
		}
	}
	return strings.Join(parts, " / ")
}
