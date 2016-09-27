package interop

type CategoryEntry struct {
	Key  rune
	Name string
}

type Category struct {
	Name  string
	Items []CategoryEntry
}

type InventoryItem interface {
	Name() string
	Category() string
	Weight() int
	Droppable() bool
}

type InventoryFoodItem interface {
	Cook()
	Cooked() bool
	Edible() bool
	SatiationValue() int
}

type InventoryProvider interface {
	Weight() int
	Count() int
	CanAssignRune() bool
	UnassignedRune() rune
	Add(item InventoryItem) rune
	Ref(key rune) InventoryItem
	HasKey(key rune) bool
	IsDroppable(key rune) bool
	Drop(key rune)
	RandomDrop() InventoryItem
	NameMap(categories ...string) map[rune]string
	ByCategory() []Category
}
