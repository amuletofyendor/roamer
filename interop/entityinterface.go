package interop

type VisibleThing interface {
	DisplayRune() rune
}

type EmplacedThing interface {
	X() int
	Y() int
	Place(x, y int)
	IsPassable() bool
}

type NamedThing interface {
	Name() string
}

type LivingThing interface {
	IsAlive() bool
	IsPlayer() bool
}

type BehavingThing interface {
	Tick(com Command, l Level)
}

type StatReporter interface {
	StatString() string
}

type Attacker interface {
	MakeAttack(currentTick, currentSubTick int) (int, int)
}

type Learner interface {
	AwardExpFor(e ExpProvider, currentTick, currentSubTick int)
}

type ExpProvider interface {
	AwardExp() int
}

type Attackable interface {
	AcceptAttack(damage, speed, currentTick, currentSubTick int)
}

type Looker interface {
	LookAt(object string, currentTick, currentSubTick int)
}

type Taker interface {
	Take(i InventoryItem, currentTick, currentSubTick int)
}

type Consumer interface {
	Consume(f InventoryFoodItem, currentTick, currentSubTick int)
}

type Wielder interface {
	Wield(w InventoryItem, currentTick, currentSubTick int)
}

type WorldObject interface {
	VisibleThing
	NamedThing
	EmplacedThing
	LivingThing
}

type Entity interface {
	WorldObject
	BehavingThing
	InventoryProvider
	HistoryProvider
	StatReporter
}
