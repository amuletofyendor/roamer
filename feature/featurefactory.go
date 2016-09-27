package feature

import (
	"../inventory"
	"../item"
	"../dice"
	"./behavior"
)

type FeatureFactory struct {
}

const (
	FeatureAnyItem int = iota
	FeatureGuidestone
	FeatureCampsiteSpit
	FeatureCampsiteBedroll
	FeatureTreasureChest
)

func (ff *FeatureFactory) MakeFeature(featureType int) Feature {
	var f Feature

	switch featureType {
	case FeatureAnyItem:
		f = ff.MakeAnyItem(nil)
	case FeatureTreasureChest:
		f = ff.MakeTreasureChest()
	case FeatureGuidestone:
		f = ff.MakeGuidestone()
	case FeatureCampsiteSpit:
		f = ff.MakeCampsiteSpit()
	case FeatureCampsiteBedroll:
		f = ff.MakeCampsiteBedroll()
	default:
		f = ff.MakeFeature(FeatureAnyItem)
	}

	return f
}

func (ff *FeatureFactory) MakeItem(i item.Item) Feature {
	return ff.MakeAnyItem(i)
}

func (ff *FeatureFactory) MakeAnyItem(i item.Item) Feature {
	var anyItem Feature
	var itemToAdd item.Item

	if i == nil {
		rf := item.RandomFactory{}
		itemToAdd = rf.Random()
	} else {
		itemToAdd = i
	}

	anyItem = Feature{itemToAdd.Name(),
		itemToAdd.Rune(),
		0,
		0,
		false,
		nil,
		inventory.Inventory{}}

	anyItem.Add(itemToAdd)

	return anyItem
}

func (ff *FeatureFactory) MakeTreasureChest() Feature {
	chest := Feature{"treasure chest",
		'(',
		0,
		0,
		true,
		nil,
		inventory.Inventory{}}

	for i := 1; i < dice.Oned4(); i++ {
		rf := item.RandomFactory{}
		chest.Add(rf.Random())
	}

	return chest
}

func (ff *FeatureFactory) MakeGuidestone() Feature {
	return Feature{"guidestone",
		'~',
		0,
		0,
		true,
		nil,
		inventory.Inventory{}}
}

func (ff *FeatureFactory) MakeCampsiteSpit() Feature {
	return Feature{"campsite spit",
		'A',
		0,
		0,
		true,
		nil,
		inventory.Inventory{}}
}

func (ff *FeatureFactory) MakeCampsiteBedroll() Feature {
	return Feature{"campsite bedroll",
		'=',
		0,
		0,
		true,
		behavior.CampsiteBedrollBehavior{},
		inventory.Inventory{}}
}
