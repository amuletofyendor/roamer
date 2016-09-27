package feature

import (
	"../creature"
	"../interop"
	"../inventory"
	"./behavior"
)

type Feature struct {
	name            string
	appearance      rune
	x               int
	y               int
	validIfEmpty    bool
	ambientBehavior behavior.FeatureBehavior
	inventory.Inventory
}

func (f *Feature) X() int {
	return f.x
}

func (f *Feature) Y() int {
	return f.y
}

func (f *Feature) Place(x, y int) {
	f.x = x
	f.y = y
}

func (f *Feature) Valid() bool {
	if f.validIfEmpty == false {
		return f.Count() > 0
	}

	return true
}

func (f *Feature) Name() string {
	return f.name
}

func (f *Feature) DisplayRune() rune {
	return f.appearance
}

func (f *Feature) Tick(com interop.Command, l interop.Level) {
	if f.ambientBehavior != nil {
		entities := l.EntitiesByLocation(f.x, f.y, f)

		for _, e := range entities {
			if eL, ok := e.(*creature.Creature); ok == true {
				f.ambientBehavior.Tick(eL, l.CurrentTick(), l.CurrentSubTick())
			}
		}
	}
}

func (f *Feature) IsPassable() bool {
	return true
}

func (f *Feature) IsPlayer() bool {
	return false
}

func (f *Feature) IsAlive() bool {
	return true
}

func (f *Feature) HistoryAvailable() bool {
	return false
}

func (f *Feature) HistoryViewed() {
}

func (f *Feature) History() []interop.HistoryItem {
	return nil
}

func (f *Feature) AppendHistory(event string, currentTick, currentSubTick int) {
}

func (f *Feature) StatString() string {
	return ""
}
