package inventory

import (
	"../interop"
)

type Inventory struct {
	capacity int
	itemList []interop.InventoryItem
	items    map[rune]interop.InventoryItem
}

func (iv *Inventory) Weight() int {
	total := 0
	for _, item := range iv.itemList {
		total += item.Weight()
	}
	return total
}

func (iv *Inventory) Count() int {
	return len(iv.itemList)
}

func (iv *Inventory) CanAssignRune() bool {
	return len(iv.items) < 66
}

func (iv *Inventory) UnassignedRune() rune {
	for _, testRune := range InventoryRunes {
		_, present := iv.items[testRune]
		if !present {
			return testRune
		}
	}

	return InventoryNullRune
}

func (iv *Inventory) Add(item interop.InventoryItem) rune {
	if !iv.CanAssignRune() ||
		(iv.capacity > 0 &&
			(item.Weight()+iv.Weight()) > iv.capacity) {
		return InventoryNullRune
	}

	newRune := iv.UnassignedRune()

	if iv.items == nil {
		iv.items = make(map[rune]interop.InventoryItem)
	}

	iv.itemList = append(iv.itemList, item)
	iv.items[newRune] = iv.itemList[len(iv.itemList)-1]

	return newRune
}

func (iv *Inventory) Ref(key rune) interop.InventoryItem {
	return iv.items[key]
}

func (iv *Inventory) HasKey(key rune) bool {
	_, found := iv.items[key]
	return found
}

func (iv *Inventory) IsDroppable(key rune) bool {
	return iv.HasKey(key) && iv.items[key].Droppable()
}

func (iv *Inventory) Drop(key rune) {
	item := iv.items[key]

	if item.Droppable() == true {
		delete(iv.items, key)

		for i := len(iv.itemList) - 1; i >= 0; i-- {
			if iv.itemList[i] == item {
				iv.itemList = append(iv.itemList[:i], iv.itemList[i+1:]...)
				break
			}
		}
	}
}

func (iv *Inventory) RandomDrop() interop.InventoryItem {
	for k, v := range iv.items {
		if v.Droppable() {
			iv.Drop(k)
			return v
		}
	}

	return nil
}

func (iv *Inventory) NameMap(categories ...string) map[rune]string {
	nameMap := make(map[rune]string)

	for k, v := range iv.items {
		typeOk := true

		if len(categories) > 0 {
			typeOk = false

			for _, t := range categories {
				if v.Category() == t {
					typeOk = true
					break
				}
			}
		}

		if typeOk == true {
			nameMap[k] = v.Name()
		}
	}

	return nameMap
}
