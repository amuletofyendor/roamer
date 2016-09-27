package item

import (
	"../dice"
)

type RandomFactory struct {
}

func (rf *RandomFactory) Random() Item {
	if dice.Oned6() == 6 {
		wf := WeaponFactory{}
		return wf.Random(WeaponClassAll)
	} else if dice.Oned6() == 6 {
		af := ArmorFactory{}
		return af.Random()
	} else if dice.Oned4() == 4 {
		cf := CorpseFactory{}
		return cf.Random()
	} else if dice.CoinToss() == dice.Heads {
		pf := PotionFactory{}
		return pf.Random()
	} else {
		tf := TrashFactory{}
		return tf.Random()
	}
}
