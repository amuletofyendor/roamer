package interop

type TwoDimensionalSpace interface {
  Width() int
  Height() int
  SetTile(x, y int, tile rune)
  GetTile(x, y int) rune
  TerrainPassable(x, y int) bool
}

type TimeProvider interface {
  CurrentTick() int
  CurrentSubTick() int
}

type TimeReceiver interface {
  SetCurrentTick(int)
  SetCurrentSubTick(int)
}

type PlayerContainer interface {
  Player() WorldObject
  TakePlayer() WorldObject
  PlacePlayer(player WorldObject, fromAbove bool)
}

type PopulatedSpace interface {
  Features() []WorldObject
  Creatures() []WorldObject
  AddWorldObject(wo WorldObject)
  RemoveWorldObject(wo WorldObject)
  AddFeatureByTypeAndPosition(featureType int, x, y int)
  EntitiesByLocation(x, y int, without ...WorldObject) []WorldObject
}

type ShroudedSpace interface {
  Shrouded(x, y int) bool
}

type MineableSpace interface {
  Mine(x, y int)
}

type Level interface {
  Init(universeSeed, depth uint64)
  Tick(com Command)
  Report(xStart, yStart, maxX, maxY int)
	TwoDimensionalSpace
  PlayerContainer
  TimeProvider
  TimeReceiver
  PopulatedSpace
  MineableSpace
  ShroudedSpace
}
