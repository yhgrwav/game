package entity

// FeatureKind - особенность монстра
type FeatureKind uint8

const (
	FeatureNone FeatureKind = iota
	SkeletonWeakCrushing
	SlimeImmuneChopping
	GhostHiddenAttack
	GolemStoneSkin
	DragonBreath
)

// Enemy - монстр, с которым сражается герой
type Enemy struct {
	Name    string
	HP      uint8
	Damage  uint8
	Stats   Attributes
	Feature FeatureKind
	Reward  Weapon
}
