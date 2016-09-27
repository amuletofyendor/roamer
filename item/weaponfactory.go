package item

import (
	"math/rand"
)

type WeaponFactory struct {
}

const (
	WeaponClassAll int = iota
	WeaponClassPeasant
	WeaponClassFellCreature
	WeaponClassCreaturePart
	WeaponClassKnight
)

const (
	LetterOpener int = iota
	WoodenSpoon
	Scissors
	ShortSword
	LongSword
	Mace
	Exkillibur
	GoblinDirk
	SharpFemur
	HipboneAxe
	SkullClub
	GoblinClaws
	RatClaws
	LastWeapon = RatClaws
)

func (wf *WeaponFactory) Wieldable(class int) []Weapon {
	weapons := []Weapon(nil)

	for i := 0; i <= LastWeapon; i++ {
		w := wf.MakeWeapon(i)
		if w.Droppable() &&
			(class == WeaponClassAll || w.Class() == class) {
			weapons = append(weapons, w)
		}
	}

	return weapons
}

func (wf *WeaponFactory) MakeWeapon(weaponType int) Weapon {
	switch weaponType {
	case WoodenSpoon:
		return Weapon{"wooden spoon", WeaponClassPeasant, 2, 1, true}
	case Scissors:
		return Weapon{"scissors", WeaponClassPeasant, 3, 1, true}
	case LetterOpener:
		return Weapon{"letter opener", WeaponClassPeasant, 7, 2, true}
	case GoblinDirk:
		return Weapon{"goblin dirk", WeaponClassFellCreature, 8, 3, true}
	case SharpFemur:
		return Weapon{"sharpened femur", WeaponClassFellCreature, 7, 3, true}
	case HipboneAxe:
		return Weapon{"hipbone axe", WeaponClassFellCreature, 5, 4, true}
	case SkullClub:
		return Weapon{"skull club", WeaponClassFellCreature, 6, 3, true}
	case GoblinClaws:
		return Weapon{"wicked claws", WeaponClassCreaturePart, 4, 0, false}
	case RatClaws:
		return Weapon{"tooth and claw", WeaponClassCreaturePart, 3, 0, false}
	case ShortSword:
		return Weapon{"shortsword", WeaponClassKnight, 11, 3, true}
	case LongSword:
		return Weapon{"longsword", WeaponClassKnight, 14, 4, true}
	case Mace:
		return Weapon{"mace", WeaponClassKnight, 18, 6, true}
	case Exkillibur:
		return Weapon{"exkillibur", WeaponClassKnight, 75, 4, true}
	default:
		return wf.MakeWeapon(WoodenSpoon)
	}
}

func (wf *WeaponFactory) Random(class int) Weapon {
	wieldable := wf.Wieldable(class)
	return wieldable[rand.Intn(len(wieldable))]
}
