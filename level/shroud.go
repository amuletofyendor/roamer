package level

import (
  "../util"
  "math"
)

func (l *Level) initShroud() {
  l.shroud = make([][]byte, l.height)
  for y := 0; y < l.height; y++ {
    l.shroud[y] = make([]byte, l.width)
    for x := 0; x < l.width; x++ {
      l.shroud[y][x] = 0
    }
  }
}

func (l *Level) deshroud(x, y int, withRay bool) {
  l.setShroud(x, y, 255)
  l.setShroud(x+1, y, 255)
  l.setShroud(x+2, y, 255)
  l.setShroud(x-1, y, 255)
  l.setShroud(x-2, y, 255)
  l.setShroud(x, y+1, 255)
  l.setShroud(x, y+2, 255)
  l.setShroud(x, y-1, 255)
  l.setShroud(x, y-2, 255)
  l.setShroud(x-1, y-1, 200)
  l.setShroud(x+1, y-1, 200)
  l.setShroud(x-1, y+1, 200)
  l.setShroud(x+1, y+1, 200)

  if withRay == true {
    l.deshroudByRay(x, y)
  }
}

func (l *Level) deshroudByRay(x, y int) {
  l.setShroud(x, y, 255)

  twoPi := math.Pi * 2.0
  for ray := 0; ray < 255; ray++ {
    rayRads := twoPi * (255.0 / float64(ray))
    xRatio := math.Cos(rayRads)
    yRatio := math.Sin(rayRads)
    for i := 0.5; i <= 20.0; i += 0.5 {
      rayX, rayY := x+util.Round(i*xRatio), y+util.Round(i*yRatio)
      if IsWall(l.GetTile(rayX, rayY)) || (i >= 20.0) {
        l.shroudLineDraw(x, y, rayX, rayY, 128)
        break
      }
    }
  }
}

func (l *Level) shroudLineDraw(x1, y1, x2, y2 int, value byte) {
  xOrig, yOrig := x1, y1
  dX := x2 - x1
  dY := y2 - y1

  if dX == 0 {
    l.shroudLineVert(x1, y1, y2, value)
    return
  }

  if dY == 0 {
    l.shroudLineHoriz(x1, x2, y1, value)
    return
  }

  oct := util.OctantFromDeltas(dX, dY)
  x1, y1 = util.SwitchToOctantZeroFrom(oct, x1, y1)
  x2, y2 = util.SwitchToOctantZeroFrom(oct, x2, y2)
  dX = x2 - x1
  dY = y2 - y1

  d := 2*dY - dX
  l.setShroud(xOrig, yOrig, value)
  y := 0

  if d > 0 {
    y += 1
    d = d - (2 * dX)
  }

  for x := 0; x < dX; x++ {
    drawX, drawY := util.SwitchFromOctantZeroTo(oct, x, y)
    l.setShroud(xOrig+drawX, yOrig+drawY, value)
    d = d + (2 * dY)
    if d > 0 {
      y += 1
      d = d - (2 * dX)
    }
  }
}

func (l *Level) shroudLineHoriz(x1, x2, y int, value byte) {
  xStart := util.Min(x1, x2)
  xEnd := util.Max(x1, x2)

  for x := xStart; x <= xEnd; x++ {
    l.setShroud(x, y, value)
  }
}

func (l *Level) shroudLineVert(x, y1, y2 int, value byte) {
  yStart := util.Min(y1, y2)
  yEnd := util.Max(y1, y2)

  for y := yStart; y <= yEnd; y++ {
    l.setShroud(x, y, value)
  }
}

func (l *Level) Shrouded(x, y int) bool {
  return l.getShroud(x, y) < 128
}

func (l *Level) setShroud(x, y int, value byte) {
  if x >= 0 && y >= 0 && x < l.width && y < l.height {
    l.shroud[y][x] = value
  }
}

func (l *Level) getShroud(x, y int) byte {
  if x >= 0 && y >= 0 && x < l.width && y < l.height {
    return l.shroud[y][x]
  } else {
    return 128
  }
}

func (l *Level) shroudSettle() {
  for i := 0; i < 5; i++ {
    l.shroudTick()
  }
}

func (l *Level) shroudTick() {
  for y, _ := range l.shroud {
    for x, s := range l.shroud[y] {
      if s < 128 {
        if int(l.getShroud(x-1, y))+
          int(l.getShroud(x+1, y))+
          int(l.getShroud(x, y-1))+
          int(l.getShroud(x, y+1)) >= 384 {
          l.setShroud(x, y, 128)
        }
      }
    }
  }
}
