package creature

const (
	CreaturePlayer int = iota
	CreatureGoblin
	CreatureRat
)

type CreatureBaseStats struct {
	species       int
	appearance    rune
	maxHp         int
	maxStr        int
	maxVit        int
	maxSpd        int
	expMultiplier float64
	vitMultiplier float64
	spdMultiplier float64
	metabolism    int
}
