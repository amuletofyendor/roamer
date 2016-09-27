package inventory

import (
	"../interop"
)

func (iv *Inventory) ByCategory() []interop.Category {
	categoryBuilder := make(map[string]*interop.Category)
	var categories []interop.Category

	for key, item := range iv.items {
		if _, found := categoryBuilder[item.Category()]; found == false {
			categories = append(categories, interop.Category{item.Category(), nil})
			categoryBuilder[item.Category()] = &categories[len(categories)-1]
		}

		categoryItemList := categoryBuilder[item.Category()]
		categoryItemList.Items =
			append(categoryItemList.Items, interop.CategoryEntry{key, item.Name()})
	}

	return categories
}
