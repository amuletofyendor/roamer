package item

type Potion struct {
	name      string
	weight    int
	droppable bool
}

func (p Potion) Name() string {
	return p.name
}

func (p Potion) Rune() rune {
	return '!'
}

func (p Potion) Weight() int {
	return p.weight
}

func (p Potion) Droppable() bool {
	return p.droppable
}

func (p Potion) Category() string {
	return "potions"
}
