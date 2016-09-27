package item

import (
	"fmt"
)

type Corpse struct {
	flavor string
	cooked bool
	burned bool
}

func (c Corpse) Name() string {
	cookedState := "raw"

	if c.cooked == true {
		cookedState = "cooked"
	}

	if c.burned == true {
		cookedState = "charred"
	}

	return fmt.Sprintf("%v carcass of a %v", cookedState, c.flavor)
}

func (c Corpse) Rune() rune {
	return '%'
}

func (c Corpse) Weight() int {
	return 1
}

func (c Corpse) Droppable() bool {
	return true
}

func (c Corpse) Category() string {
	return "edibles"
}

func (c Corpse) Cook() {
	if c.cooked {
		c.burned = true
	} else {
		c.cooked = true
	}
}

func (c Corpse) Cooked() bool {
	return c.cooked
}

func (c Corpse) Edible() bool {
	return !c.burned
}

func (c Corpse) SatiationValue() int {
	if c.Edible() {
		if c.Cooked() {
			return 22
		} else {
			return 15
		}
	} else {
		return 0
	}
}
