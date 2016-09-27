package item

type Armor struct {
	name       string
	protection int
	weight     int
	droppable  bool
}

func (a Armor) Name() string {
	return a.name
}

func (a Armor) Weight() int {
	return a.weight
}

func (a Armor) Protection() int {
	return a.protection
}

func (a Armor) Droppable() bool {
	return a.droppable
}

func (a Armor) Rune() rune {
	return '['
}

func (a Armor) Category() string {
	return "armor"
}
