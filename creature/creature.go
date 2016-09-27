package creature

import (
	"../interop"
	"../inventory"
	"./behavior"
)

type Creature struct {
	name           string
	x              int
	y              int
	behavior       behavior.CreatureBehavior
	baseStats      CreatureBaseStats
	level          int
	exp            int
	hp             int
	history        []interop.HistoryItem
	historySeen    bool
	equippedWeapon interop.InventoryItem
	vitBoost       int
	satiation      int
	inventory.Inventory
}

func (c *Creature) X() int {
	return c.x
}

func (c *Creature) Y() int {
	return c.y
}

func (c *Creature) Place(x, y int) {
	c.x = x
	c.y = y
}

func (c *Creature) Name() string {
	return c.name
}

func (c *Creature) DisplayRune() rune {
	return c.baseStats.appearance
}

func (c *Creature) Tick(com interop.Command, l interop.Level) {
	if c.behavior != nil {
		c.behavior.Tick(c, com, l)
		c.StatTick(l.CurrentTick(), l.CurrentSubTick()+1)
	}
}

func (c *Creature) IsPassable() bool {
	return false
}

func (c *Creature) IsPlayer() bool {
	return c.baseStats.species == CreaturePlayer
}

func (c *Creature) History() []interop.HistoryItem {
	return c.history
}

func (c *Creature) StatString() string {
	return c.AsString()
}

func (c *Creature) BoostVitality(vitBoost int) {
	c.vitBoost = vitBoost
}
