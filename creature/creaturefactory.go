package creature

import (
	"../inventory"
	"../item"
	"../dice"
	"./behavior"
)

type CreatureFactory struct {
}

func (cf *CreatureFactory) MakeCreatureBaseStats(species int) CreatureBaseStats {
	switch species {
	case CreaturePlayer:
		return CreatureBaseStats{CreaturePlayer, '@', 100, 10, 80, 80, 1000.0, 12, 12, 85}
	case CreatureGoblin:
		return CreatureBaseStats{CreatureGoblin, 'g', 50, 8, 25, 25, 100.0, 5, 5, 0}
	case CreatureRat:
		return CreatureBaseStats{CreatureRat, 'r', 20, 2, 50, 80, 20.0, 5, 20, 0}
	default:
		return cf.MakeCreatureBaseStats(CreatureRat)
	}
}

func (cf *CreatureFactory) NameCreature(species int) string {
	switch species {
	case CreaturePlayer:
		return "player"
	case CreatureGoblin:
		return "goblin"
	case CreatureRat:
		return "rat"
	default:
		return cf.NameCreature(CreatureRat)
	}
}

func (cf *CreatureFactory) BehaviorFor(species int) behavior.CreatureBehavior {
	switch species {
	case CreaturePlayer:
		return behavior.PlayerBehavior{}
	default:
		return behavior.FollowBehavior{}
	}
}

func (cf *CreatureFactory) WeaponFor(species, level int) item.Weapon {
	wf := item.WeaponFactory{}

	switch species {
	case CreaturePlayer:
		return wf.Random(item.WeaponClassPeasant)
	case CreatureGoblin:
		if dice.Oned4() == 4 {
			return wf.Random(item.WeaponClassFellCreature)
		} else {
			return wf.MakeWeapon(item.GoblinClaws)
		}
	case CreatureRat:
		return wf.MakeWeapon(item.RatClaws)
	default:
		return cf.WeaponFor(CreatureRat, level)
	}
}

func (cf *CreatureFactory) MakeCreature(species, level int) Creature {
	iv := inventory.Inventory{}
	weaponRune := iv.Add(cf.WeaponFor(species, level))

	creature := Creature{cf.NameCreature(species),
		0,
		0,
		cf.BehaviorFor(species),
		cf.MakeCreatureBaseStats(species),
		level,
		1000 * (level - 1),
		0,
		nil,
		true,
		iv.Ref(weaponRune),
		0,
		75,
		inventory.Inventory{}}

	creature.Init()

	return creature
}
