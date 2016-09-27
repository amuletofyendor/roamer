package item

type Weapon struct {
	name       string
	class      int
	baseDamage int
	weight     int
	droppable  bool
}

func (w Weapon) Name() string {
	return w.name
}

func (w Weapon) Rune() rune {
	return ')'
}

func (w Weapon) Weight() int {
	return w.weight
}

func (w Weapon) BaseDamage() int {
	return w.baseDamage
}

func (w Weapon) Class() int {
	return w.class
}

func (w Weapon) Droppable() bool {
	return w.droppable
}

func (c Weapon) Category() string {
	return "weapons"
}
