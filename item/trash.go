package item

type Trash struct {
	name      string
	weight    int
	droppable bool
}

func (t Trash) Name() string {
	return t.name
}

func (t Trash) Rune() rune {
	return '*'
}

func (t Trash) Weight() int {
	return t.weight
}

func (t Trash) Droppable() bool {
	return t.droppable
}

func (t Trash) Category() string {
	return "misc"
}
