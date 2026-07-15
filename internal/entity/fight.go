package entity

// FightContext - доступ к участникам удара и номеру хода для замыканий бонусов.
type FightContext struct {
	Attacker *Combatant
	Defender *Combatant
	Round    uint8
}
