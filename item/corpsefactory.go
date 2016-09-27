package item

import (
	"math/rand"
)

type CorpseFactory struct {
}

func (cf *CorpseFactory) MakeCorpse(name string) Corpse {
	return Corpse{name, false, false}
}

func (cf *CorpseFactory) Random() Corpse {
	return cf.MakeCorpse([]string{"rat", "dog", "cat", "goblin"}[rand.Intn(4)])
}
