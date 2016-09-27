package creature

import (
	"../history"
	"fmt"
)

const (
	Attacking              = "lashes out!"
	AttackingWith          = "attacks with %s!"
	Attacked               = "was attacked!"
	Evaded                 = "dodged the attack!"
	Healed                 = "healed %d HP."
	TookDamage             = "was hit for %d HP!"
	TookNoDamage           = "took no damage whatsoever!"
	Died                   = "died."
	GainedExp              = "gained %d exp!"
	GainedLevels           = "levelled up by a whopping %d levels!"
	GainedALevel           = "levelled up!"
	LookAtSomething        = "sees %s here"
	TookSomething          = "took %s"
	EquippedSomething      = "equipped %s"
	KnowsTheWay            = "feels the guidestone urging them forward"
	PlaceToRest            = "takes a moment to rest on a bedroll"
	HungerPreventedHealing = "is too hungry to heal naturally"
	AteSomething           = "thought that was tasty!"
	GastricDistress        = "experienced sudden gastric distress!"
)

func (c *Creature) HistoryAvailable() bool {
	return c.historySeen == false
}

func (c *Creature) HistoryViewed() {
	c.history = nil
	c.historySeen = true
}

func (c *Creature) AppendHistory(event string, currentTick, currentSubTick int) {
	hf := history.HistoryFactory{}
	historyItem := hf.MakeHistory(event, currentTick, currentSubTick)
	c.history = append(c.history, &historyItem)
	c.historySeen = false
}

func (c *Creature) RecordExpGain(exp, currentTick, currentSubTick int) {
	c.AppendHistory(fmt.Sprintf(GainedExp, exp), currentTick, currentSubTick)
}

func (c *Creature) RecordLevelGain(delta, currentTick, currentSubTick int) {
	if delta > 1 {
		c.AppendHistory(fmt.Sprintf(GainedLevels, delta), currentTick, currentSubTick)
	} else {
		c.AppendHistory(GainedALevel, currentTick, currentSubTick)
	}
}

func (c *Creature) RecordDamage(amount, currentTick, currentSubTick int) {
	if amount > 0 {
		c.AppendHistory(fmt.Sprintf(TookDamage, amount), currentTick, currentSubTick)
	} else {
		c.AppendHistory(fmt.Sprintf(TookNoDamage), currentTick, currentSubTick)
	}
}

func (c *Creature) RecordHeal(amount, currentTick, currentSubTick int) {
	c.AppendHistory(fmt.Sprintf(Healed, amount), currentTick, currentSubTick)
}

func (c *Creature) RecordEvadeAttack(currentTick, currentSubTick int) {
	c.AppendHistory(Evaded, currentTick, currentSubTick)
}

func (c *Creature) RecordDeath(currentTick, currentSubTick int) {
	c.AppendHistory(Died, currentTick, currentSubTick)
}

func (c *Creature) RecordAcceptAttack(currentTick, currentSubTick int) {
}

func (c *Creature) RecordMakeAttack(weapon string, currentTick, currentSubTick int) {
	if len(weapon) == 0 {
		c.AppendHistory(Attacking, currentTick, currentSubTick)
	} else {
		c.AppendHistory(fmt.Sprintf(AttackingWith, weapon), currentTick, currentSubTick)
	}
}

func (c *Creature) RecordLookAt(objectName string, currentTick, currentSubTick int) {
	c.AppendHistory(fmt.Sprintf(LookAtSomething, objectName), currentTick, currentSubTick)
}

func (c *Creature) RecordTake(objectName string, currentTick, currentSubTick int) {
	c.AppendHistory(fmt.Sprintf(TookSomething, objectName), currentTick, currentSubTick)
}

func (c *Creature) RecordEquip(objectName string, currentTick, currentSubTick int) {
	c.AppendHistory(fmt.Sprintf(EquippedSomething, objectName), currentTick, currentSubTick)
}

func (c *Creature) RecordGuidestoneInteraction(currentTick, currentSubTick int) {
	c.AppendHistory(KnowsTheWay, currentTick, currentSubTick)
}

func (c *Creature) RecordBedrollInteraction(currentTick, currentSubTick int) {
	c.AppendHistory(PlaceToRest, currentTick, currentSubTick)
}

func (c *Creature) RecordHungerPreventedHealing(currentTick, currentSubTick int) {
	c.AppendHistory(HungerPreventedHealing, currentTick, currentSubTick)
}

func (c *Creature) RecordSatiation(currentTick, currentSubTick int) {
	c.AppendHistory(AteSomething, currentTick, currentSubTick)
}

func (c *Creature) RecordGasticDistress(currentTick, currentSubTick int) {
	c.AppendHistory(GastricDistress, currentTick, currentSubTick)
}
