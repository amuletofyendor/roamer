package interop

type Universe interface {
  SetSeed(seed uint64)
  Player() WorldObject
  Advance(com Command) bool
  Report(maxX, maxY int)
}
