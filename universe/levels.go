package universe

import (
  "../creature"
  "../interop"
  "../level"
)

func (u *Universe) CurrentLevel() interop.Level {
  if u.currentLevelIndex < uint64(len(u.levels)) {
    return u.levels[u.currentLevelIndex]
  } else {
    return nil
  }
}

func (u *Universe) LevelCount() uint64 {
  return uint64(len(u.levels))
}

func (u *Universe) GoToLevel(index uint64) {
  var player interop.WorldObject
  playerFound := false

  if u.CurrentLevel() != nil {
    player = u.CurrentLevel().TakePlayer()
    playerFound = true
  }

  goingDown := u.currentLevelIndex <= index

  u.currentLevelIndex = index

  if index >= u.LevelCount() {
    for i := u.LevelCount(); i <= index; i++ {
      u.GenerateLevel(i + 1)
    }
  }

  if !playerFound {
    ef := creature.CreatureFactory{}
    newPlayer := ef.MakeCreature(creature.CreaturePlayer, 1)
    player = &newPlayer
  }

  u.CurrentLevel().PlacePlayer(player, goingDown)
  u.CurrentLevel().SetCurrentTick(u.turn)
}

func (u *Universe) GenerateLevel(depth uint64) {
  l := level.Level{}
  l.Init(u.seed, depth)
  u.AddLevel(&l)
}

func (u *Universe) AddLevel(l interop.Level) {
  u.levels = append(u.levels, l)
}
