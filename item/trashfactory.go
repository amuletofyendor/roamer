package item

import (
	"math/rand"
)

type TrashFactory struct {
}

const (
	Gravel int = iota
	SmallStone
	WarmPebble
	GallStone
	Bezoar
	Twig
	Stick
	Twine
	Tooth
	RottenTooth
	RatSkull
	BrokenHilt
	BrokenBlade
	BrokenBladeTip
	AshPile
	SaltPile
	CharcoalPile
	BrokenCrockery
	WoodenDoll
	LeatherDoll
	CrushedEye
	Feather
	SmallBone
)

func (wf *TrashFactory) MakeTrash(trashType int) Trash {
	switch trashType {
	case Gravel:
		return Trash{"gravel", 1, true}
	case SmallStone:
		return Trash{"small stone", 1, true}
	case WarmPebble:
		return Trash{"warm pebble", 1, true}
	case GallStone:
		return Trash{"gall stone", 1, true}
	case Bezoar:
		return Trash{"beozar", 1, true}
	case Twig:
		return Trash{"gnarled twig", 1, true}
	case Stick:
		return Trash{"straight stick", 1, true}
	case Twine:
		return Trash{"piece of twine", 1, true}
	case Tooth:
		return Trash{"tooth", 1, true}
	case RottenTooth:
		return Trash{"rotten tooth", 1, true}
	case RatSkull:
		return Trash{"rat skull", 1, true}
	case BrokenHilt:
		return Trash{"broken hilt", 2, true}
	case BrokenBlade:
		return Trash{"broken blade", 2, true}
	case BrokenBladeTip:
		return Trash{"broken blade tip", 1, true}
	case AshPile:
		return Trash{"ash pile", 1, true}
	case SaltPile:
		return Trash{"salt pile", 1, true}
	case CharcoalPile:
		return Trash{"charcoal pile", 1, true}
	case BrokenCrockery:
		return Trash{"piece of broken crockery", 1, true}
	case WoodenDoll:
		return Trash{"small wooden doll", 1, true}
	case LeatherDoll:
		return Trash{"small leather doll", 1, true}
	case CrushedEye:
		return Trash{"crushed eyeball", 1, true}
	case Feather:
		return Trash{"feather", 1, true}
	case SmallBone:
		return Trash{"small bone", 1, true}
	default:
		return wf.MakeTrash(SmallStone)
	}
}

func (tf *TrashFactory) Random() Trash {
	return tf.MakeTrash(rand.Intn(SmallBone))
}
