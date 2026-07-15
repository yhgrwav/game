package entity

// Logic - модификатор урона: всё для решения бонус берёт из FightContext -
// кто бьёт, кого бьют и какой сейчас ход
type Logic func(ctx FightContext, damage uint16) uint16

type BonusName string

const (
	Rage            BonusName = "rage"
	StoneSkin       BonusName = "stone_skin"
	HiddenAttack    BonusName = "hidden_attack"
	Shield          BonusName = "shield"
	ImpulseToAction BonusName = "impulse_to_action"
	Poison          BonusName = "poison"
	AgilityUp       BonusName = "agility_up"
	StrengthUp      BonusName = "strength_up"
	StaminaUp       BonusName = "stamina_up"
)

type Bonus struct {
	Title         string
	Description   string
	Income        Logic
	Outcome       Logic
	AttributeGain func(h *Hero)
}
