package level

import (
  "../dice"
)

func (l *Level) Mine(x, y int) {
  tile := l.GetTile(x, y)
  if dice.Oned4() > 2 || !IsWall(tile) {
    return
  }

  newTile := tile

  switch tile {
  case Wall:
    newTile = DamagedWall
  case DamagedWall:
    newTile = VeryDamagedWall
  case VeryDamagedWall:
    newTile = FreeTile
  }

  l.SetTile(x, y, newTile)
}
