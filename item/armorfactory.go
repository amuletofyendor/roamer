package item

import (
	"math/rand"
)

type ArmorFactory struct {
}

const (
	PeasantCap int = iota
	TwigHelmet
	CottonShirt
	Moccasins
	RatskinCap
	CarapaceChest
	CarapaceBoots
)

func (af *ArmorFactory) MakeArmor(armorType int) Armor {
	switch armorType {
	case PeasantCap:
		return Armor{"peasant cap", 1, 1, true}
	case CottonShirt:
		return Armor{"cotton shirt", 1, 2, true}
	case Moccasins:
		return Armor{"moccasins", 1, 0, true}
	case RatskinCap:
		return Armor{"ratskin cap", 2, 1, true}
	case CarapaceChest:
		return Armor{"carapace chestpiece", 5, 5, true}
	case CarapaceBoots:
		return Armor{"carapace boots", 3, 3, true}
	case TwigHelmet:
		return Armor{"twig helmet", 1, 1, true}
	default:
		return af.MakeArmor(TwigHelmet)
	}
}

func (af *ArmorFactory) Random() Armor {
	return af.MakeArmor([]int{PeasantCap,
		CottonShirt,
		Moccasins,
		RatskinCap,
		CarapaceChest,
		CarapaceBoots}[rand.Intn(6)])
}
