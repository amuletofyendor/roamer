package item

type Item interface {
	Name() string
	Rune() rune
	Category() string
	Weight() int
	Droppable() bool
}
