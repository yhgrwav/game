package entity

// Combatant - боец в бою: единое рантайм-представление, к которому приводятся
// Hero и Enemy; Feature несёт спец монстра, не выразимый бонусом (у героя
// FeatureNone), HP - uint16 под тип урона
type Combatant struct {
	Name       string
	Attributes Attributes
	Weapon     Weapon
	Bonuses    []BonusName
	Feature    FeatureKind
	HP         uint16
}
