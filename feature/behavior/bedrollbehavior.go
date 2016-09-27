package behavior

import (
	"../../creature"
)

type CampsiteBedrollBehavior struct {
}

func (b CampsiteBedrollBehavior) Tick(c *creature.Creature,
	currentTick,
	currentSubTick int) {
	c.RecordBedrollInteraction(currentTick, currentSubTick)
	c.BoostVitality(99)
}
