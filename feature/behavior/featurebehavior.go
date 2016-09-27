package behavior

import (
	"../../creature"
)

type FeatureBehavior interface {
	Tick(c *creature.Creature, currentTick, currentSubTick int)
}
