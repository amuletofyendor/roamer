package universe

import (
  "../interop"
  "../util"
  "fmt"
)

type Universe struct {
  levels            []interop.Level
  turn              int
  quitLevel         int
  currentLevelIndex uint64
  seed              uint64
}

func (u *Universe) SetSeed(seed uint64) {
  u.seed = seed
}

func (u *Universe) Player() interop.WorldObject {
  var player interop.WorldObject

  for _, l := range u.levels {
    player = l.Player()
    if player != nil {
      break
    }
  }

  return player
}

func (u *Universe) Advance(com interop.Command) bool {
  if !u.Player().(interop.Entity).IsAlive() {
    return u.HandleDeath(com)
  } else {
    u.turn = u.turn + 1

    // if com.IsDescend() {
    //  if u.Player().X() == u.CurrentLevel().downStairsX &&
    //    u.Player().Y() == u.CurrentLevel().downStairsY {
    //    u.GoToLevel(uint64(u.currentLevelIndex + 1))
    //  }
    // } else if com.IsAscend() {
    //  if u.Player().X() == u.CurrentLevel().upStairsX &&
    //    u.Player().Y() == u.CurrentLevel().upStairsY &&
    //    u.currentLevelIndex > 0 {
    //    u.GoToLevel(uint64(u.currentLevelIndex - 1))
    //  }
    // }

    u.CurrentLevel().Tick(com)

    u.HandleQuitting(com)
    return u.ReallyQuit()
  }
}

func (u *Universe) Report(maxX, maxY int) {
  util.StringOut(0,
    0,
    fmt.Sprintf("DLvl: %d %s %s",
      u.currentLevelIndex+1,
      u.Player().(interop.Entity).Name(),
      u.Player().(interop.Entity).StatString()))

  u.CurrentLevel().Report(0, 1, maxX, maxY-2)
}
