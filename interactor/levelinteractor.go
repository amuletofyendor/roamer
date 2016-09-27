package interactor

import (
	"../interop"
	"fmt"
)

type LevelInteractor struct {
	Level       interop.Level
	Worldobject interop.WorldObject
}

func (li *LevelInteractor) Move(lt, rt, up, dn bool) {
	originalX := li.Worldobject.X()
	originalY := li.Worldobject.Y()
	newX := originalX
	newY := originalY

	if lt {
		newX -= 1
	}

	if rt {
		newX += 1
	}

	if up {
		newY -= 1
	}

	if dn {
		newY += 1
	}

	if newX < 0 {
		newX = 0
	}
	if newX >= li.Level.Width() {
		newX = li.Level.Width() - 1
	}
	if newY < 0 {
		newY = 0
	}
	if newY >= li.Level.Height() {
		newY = li.Level.Height() - 1
	}

	beings := li.Level.Creatures()
	for i, _ := range beings {
		if li.Worldobject != beings[i] &&
			!beings[i].IsPassable() &&
			newX == beings[i].X() &&
			newY == beings[i].Y() {
			newX = originalX
			newY = originalY

			attackableEntity, ok := beings[i].(interop.Attackable)

			if ok == true {
				li.Attack(attackableEntity)
				return
			}
		}
	}

	if !li.ValidMove(newX, newY) {
		li.Level.Mine(newX, newY)
		newX = originalX
		newY = originalY
		return
	}

	li.Worldobject.Place(newX, newY)
}

func (li *LevelInteractor) Attack(target interop.Attackable) {
	attacker, ok := li.Worldobject.(interop.Attacker)

	if ok == false {
		return
	}

	dmg, spd := attacker.MakeAttack(
		li.Level.CurrentTick(),
		li.Level.CurrentSubTick())

	target.AcceptAttack(
		dmg,
		spd,
		li.Level.CurrentTick(),
		li.Level.CurrentSubTick()+1)

	if attacker.(interop.LivingThing).IsAlive() &&
		!target.(interop.LivingThing).IsAlive() {
		attacker.(interop.Learner).AwardExpFor(target.(interop.ExpProvider),
			li.Level.CurrentTick(),
			li.Level.CurrentSubTick()+2)
	}
}

func (li *LevelInteractor) Look(x, y int) {
	objects := li.Level.EntitiesByLocation(x, y, li.Worldobject)

	if len(objects) > 0 {
		time := 0
		for _, obj := range objects {
			e := obj.(interop.Entity)
			li.Worldobject.(interop.Looker).LookAt(e.Name(),
				li.Level.CurrentTick(),
				li.Level.CurrentSubTick()+time)
			time++

			itemsByCategory := e.ByCategory()

			for _, category := range itemsByCategory {
				li.Worldobject.(interop.HistoryProvider).AppendHistory(fmt.Sprintf("[%s]", category.Name),
					li.Level.CurrentTick(),
					li.Level.CurrentSubTick()+time)
				time++

				for _, item := range category.Items {
					li.Worldobject.(interop.HistoryProvider).AppendHistory(fmt.Sprintf("  %c - %s", item.Key, item.Name),
						li.Level.CurrentTick(),
						li.Level.CurrentSubTick()+time)
					time++
				}
			}
		}
	} else {
		li.Worldobject.(interop.Looker).LookAt("nothing in particular",
			li.Level.CurrentTick(),
			li.Level.CurrentSubTick())
	}
}

func (li *LevelInteractor) Take(x, y int) {
	objects := li.Level.EntitiesByLocation(x, y, li.Worldobject)

	if len(objects) > 0 {
		time := 0
		for _, obj := range objects {
			e := obj.(interop.Entity)

			if e.Count() > 0 {
				nameMap := e.NameMap()

				for k, _ := range nameMap {
					if e.IsDroppable(k) {
						item := e.Ref(k)
						li.Worldobject.(interop.Taker).Take(item,
							li.Level.CurrentTick(),
							li.Level.CurrentSubTick()+time)
						e.Drop(k)
						time++
					}
				}
			}
		}
	}
}

func (li *LevelInteractor) Eat(inventoryKey rune) {
	food := li.Worldobject.(interop.Entity).Ref(inventoryKey)
	li.Worldobject.(interop.Entity).Drop(inventoryKey)
	li.Worldobject.(interop.Consumer).Consume(food.(interop.InventoryFoodItem),
		li.Level.CurrentTick(),
		li.Level.CurrentSubTick())
}

func (li *LevelInteractor) Wield(inventoryKey rune) {
	weapon := li.Worldobject.(interop.Entity).Ref(inventoryKey)
	li.Worldobject.(interop.Wielder).Wield(weapon,
		li.Level.CurrentTick(),
		li.Level.CurrentSubTick())
}

func (li *LevelInteractor) ValidMove(x, y int) bool {
	if !li.Level.TerrainPassable(x, y) {
		return false
	}

	for _, o := range li.Level.EntitiesByLocation(x, y, li.Worldobject) {
		if !o.IsPassable() &&
			o.X() == x &&
			o.Y() == y {
			return false
		}
	}

	return true
}
