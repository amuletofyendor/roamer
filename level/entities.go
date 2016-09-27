package level

import (
  "../creature"
  "../feature"
  "../interop"
  "../item"
  "../dice"
  "math/rand"
)

func (l *Level) Player() interop.WorldObject {
  var player interop.WorldObject

  for i, _ := range l.entities {
    if l.entities[i].IsPlayer() {
      player = l.entities[i]
    }
  }

  return player
}

func (l *Level) Features() []interop.WorldObject {
  var list []interop.WorldObject

  for _, e := range l.entities {
    if _, ok := e.(*feature.Feature); ok == true {
      list = append(list, e)
    }
  }

  return list
}

func (l *Level) Creatures() []interop.WorldObject {
  var list []interop.WorldObject

  for _, e := range l.entities {
    if _, ok := e.(*creature.Creature); ok == true {
      list = append(list, e)
    }
  }

  return list
}

func (l *Level) renderList() *[]interop.WorldObject {
  if l.entityRenderList == nil {
    l.entityRenderList = append(l.entityRenderList, l.Features()...)
    l.entityRenderList = append(l.entityRenderList, l.Creatures()...)
  }

  return &l.entityRenderList
}

func (l *Level) TakePlayer() interop.WorldObject {
  p := l.Player()
  l.RemoveWorldObject(p)
  return p
}

func (l *Level) PlacePlayer(player interop.WorldObject, fromAbove bool) {
  if fromAbove == true {
    player.Place(l.upStairsX, l.upStairsY)
  } else {
    player.Place(l.downStairsX, l.downStairsY)
  }

  l.AddWorldObject(player)
  l.deshroudByRay(l.upStairsX, l.upStairsY)
  l.shroudSettle()
}

func (l *Level) AddWorldObject(wo interop.WorldObject) {
  l.entities = append(l.entities, wo.(interop.Entity))
  l.entityRenderList = nil
}

func (l *Level) RemoveWorldObject(wo interop.WorldObject) {
  index := -1

  for i, _ := range l.entities {
    if l.entities[i] == wo {
      index = i
      break
    }
  }

  if index != -1 {
    l.entities = append(l.entities[:index], l.entities[index+1:]...)
    l.entityRenderList = nil
  }
}

func (l *Level) Corpseify(e *creature.Creature) {
  // This should belong to levelInteractor
  fef := feature.FeatureFactory{}

  cf := item.CorpseFactory{}
  corpsePile := fef.MakeItem(cf.MakeCorpse(e.Name()))
  corpsePile.Place(e.X(), e.Y())
  l.AddWorldObject(&corpsePile)

  item := e.RandomDrop()

  if item != nil {
    corpsePile.Add(item)
  }

  l.RemoveWorldObject(e)
}

func (l *Level) AddFeatureByTypeAndPosition(featureType int, x, y int) {
  fef := feature.FeatureFactory{}
  f := fef.MakeFeature(featureType)
  f.Place(x, y)
  l.AddWorldObject(&f)
}

func (l *Level) EntitiesByLocation(x, y int,
  without ...interop.WorldObject) []interop.WorldObject {
  var entities []interop.WorldObject

  for i, _ := range l.entities {
    e := l.entities[i]

    ignoreThisEntity := false

    for _, ignore := range without {
      if e == ignore {
        ignoreThisEntity = true
        break
      }
    }

    if !ignoreThisEntity &&
      (e.X() == x && e.Y() == y) {
      entities = append(entities, e)
    }
  }

  return entities
}

func (l *Level) tidyUp() {
  for i := len(l.entities) - 1; i >= 0; i-- {
    if f, ok := l.entities[i].(*feature.Feature); ok == true {
      if f.Valid() == false {
        l.RemoveWorldObject(l.entities[i])
      }
    }
  }
}

func (l *Level) populate() {
  ef := creature.CreatureFactory{}

  goblinCount := dice.Dice{4, 4, 10}.Roll()
  for i := 0; i < goblinCount; i++ {
    e := ef.MakeCreature(creature.CreatureGoblin, int(l.depth))
    e.Place(rand.Intn(l.width), rand.Intn(l.height))
    l.AddWorldObject(&e)
  }

  ratCount := dice.Dice{4, 8, 10}.Roll()
  for i := 0; i < ratCount; i++ {
    e := ef.MakeCreature(creature.CreatureRat, int(l.depth))
    e.Place(rand.Intn(l.width), rand.Intn(l.height))
    l.AddWorldObject(&e)
  }
}
