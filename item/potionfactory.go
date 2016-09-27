package item

import (
	"math/rand"
)

type PotionFactory struct {
}

const (
	Water int = iota
	Healing
)

func (pf *PotionFactory) MakePotion(potionType int) Potion {
	switch potionType {
	case Water:
		return Potion{"potion of water", 1, true}
	case Healing:
		return Potion{"potion of healing", 1, true}
	default:
		return pf.MakePotion(Water)
	}
}

func (pf *PotionFactory) Random() Potion {
	return pf.MakePotion([]int{Water, Healing}[rand.Intn(2)])
}
