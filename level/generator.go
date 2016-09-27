package level

import (
  "../feature"
  "../util"
  "../dice"
  "math"
  "math/rand"
)

const (
  LevelMaxX = 500
  LevelMaxY = 500
)

const (
  FreeTile         = ' '
  InaccessibleTile = '\u2591'
  Wall             = '\u2588'
  DamagedWall      = '\u2593'
  VeryDamagedWall  = '\u2592'
  Upstairs         = '<'
  Downstairs       = '>'
)

func (l *Level) initGeography() {
  l.width = 50 + rand.Intn(LevelMaxX-50)
  l.height = 50 + rand.Intn(LevelMaxY-50)

  l.layout = make([][]rune, l.height)
  for i, _ := range l.layout {
    l.layout[i] = make([]rune, l.width)
    for j, _ := range l.layout[i] {
      l.layout[i][j] = FreeTile
    }
  }

  rand.Seed(int64(l.seed + l.depth))

  l.randomWalls(15)
  l.growWalls(5+util.Minint64(l.depth, 3), 30)
  l.placeCampsites(dice.Oned4())
  l.placeDownstairs(dice.Oned20() + 5)
  l.placeUpstairs(dice.Oned20() + 5)
  l.placeGuidestones()
  l.placeItems()
  l.initShroud()
}

func IsWall(tile rune) bool {
  return tile == Wall ||
    tile == DamagedWall ||
    tile == VeryDamagedWall
}

func randomWallType() rune {
  if dice.Oned10() == 10 {
    return VeryDamagedWall
  } else if dice.Oned10() > 9 {
    return DamagedWall
  } else {
    return Wall
  }
}

func (l *Level) TerrainPassable(x, y int) bool {
  tile := l.GetTile(x, y)
  return tile == FreeTile ||
    tile == Upstairs ||
    tile == Downstairs
}

func (l *Level) randomWalls(density int) {
  for y, row := range l.layout {
    for x, _ := range row {
      l.SetTile(x, y, FreeTile)

      if rand.Intn(density) == 1 {
        l.SetTile(x, y, randomWallType())

        if dice.CoinToss() == dice.Heads {
          l.SetTile(x+1, y, randomWallType())
        }

        if dice.CoinToss() == dice.Heads {
          l.SetTile(x-1, y, randomWallType())
        }

        if dice.CoinToss() == dice.Heads {
          l.SetTile(x, y+1, randomWallType())
        }

        if dice.CoinToss() == dice.Heads {
          l.SetTile(x, y-1, randomWallType())
        }
      }
    }
  }
}

func (l *Level) placeItems() {
  itemCount := 2 + rand.Intn(175)

  for i := 0; i < itemCount; i++ {
    if dice.Oned10() == 10 {
      l.AddFeatureByTypeAndPosition(feature.FeatureTreasureChest,
        rand.Intn(l.width),
        rand.Intn(l.height))
    } else {
      l.AddFeatureByTypeAndPosition(feature.FeatureAnyItem,
        rand.Intn(l.width),
        rand.Intn(l.height))
    }
  }
}

func (l *Level) growWalls(cycles, newGrowthDensity uint64) {
  for i := uint64(0); i < cycles; i++ {
    for y, row := range l.layout {
      for x, _ := range row {
        if l.GetTile(x, y) == FreeTile {
          if (l.GetTile(x-1, y) != FreeTile &&
            l.GetTile(x+1, y) != FreeTile) ||
            (l.GetTile(x, y-1) != FreeTile &&
              l.GetTile(x, y+1) != FreeTile) ||
            (l.GetTile(x-1, y-1) != FreeTile &&
              l.GetTile(x+1, y+1) != FreeTile) ||
            (l.GetTile(x-1, y+1) != FreeTile &&
              l.GetTile(x+1, y-1) != FreeTile) {
            l.SetTile(x, y, DamagedWall)
          } else if (rand.Int63n(int64(newGrowthDensity)) == 1) &&
            (l.GetTile(x-1, y) != FreeTile ||
              l.GetTile(x+1, y) != FreeTile ||
              l.GetTile(x, y-1) != FreeTile ||
              l.GetTile(x, y+1) != FreeTile ||
              l.GetTile(x-1, y-1) != FreeTile ||
              l.GetTile(x-1, y+1) != FreeTile ||
              l.GetTile(x+1, y-1) != FreeTile ||
              l.GetTile(x+1, y+1) != FreeTile) {
            l.SetTile(x, y, VeryDamagedWall)
          }
        }
      }
    }
  }
}

func (l *Level) placeDownstairs(clearRadius int) {
  l.downStairsX = clearRadius + rand.Intn(l.width-clearRadius)
  l.downStairsY = clearRadius + rand.Intn(l.height-clearRadius)
  l.clearRadius(l.downStairsX, l.downStairsY, clearRadius)
  l.SetTile(l.downStairsX, l.downStairsY, Downstairs)
  l.SetTile(l.downStairsX-1, l.downStairsY+1, Wall)
  l.SetTile(l.downStairsX, l.downStairsY+1, Wall)
  l.SetTile(l.downStairsX+1, l.downStairsY+1, Wall)
  l.SetTile(l.downStairsX-1, l.downStairsY, DamagedWall)
  l.SetTile(l.downStairsX+1, l.downStairsY, DamagedWall)
  l.SetTile(l.downStairsX-1, l.downStairsY-1, VeryDamagedWall)
  l.SetTile(l.downStairsX+1, l.downStairsY-1, VeryDamagedWall)
}

func (l *Level) placeUpstairs(clearRadius int) {
  l.upStairsX = clearRadius + rand.Intn(l.width-clearRadius)
  l.upStairsY = clearRadius + rand.Intn(l.height-clearRadius)
  l.clearRadius(l.upStairsX, l.upStairsY, clearRadius)
  l.SetTile(l.upStairsX, l.upStairsY, Upstairs)
  l.SetTile(l.upStairsX-1, l.upStairsY-1, Wall)
  l.SetTile(l.upStairsX, l.upStairsY-1, Wall)
  l.SetTile(l.upStairsX+1, l.upStairsY-1, Wall)
  l.SetTile(l.upStairsX-1, l.upStairsY, DamagedWall)
  l.SetTile(l.upStairsX+1, l.upStairsY, DamagedWall)
  l.SetTile(l.upStairsX-1, l.upStairsY+1, VeryDamagedWall)
  l.SetTile(l.upStairsX+1, l.upStairsY+1, VeryDamagedWall)
}

func (l *Level) placeGuidestones() {
  stairsXDiff := float64(l.downStairsX - l.upStairsX)
  stairsYDiff := float64(l.downStairsY - l.upStairsY)
  ang := math.Atan(stairsYDiff / stairsXDiff)

  cosAng := math.Cos(ang)
  sinAng := math.Sin(ang)

  l.AddFeatureByTypeAndPosition(feature.FeatureGuidestone,
    l.upStairsX+int(math.Copysign(10*cosAng, stairsXDiff)),
    l.upStairsY+int(math.Copysign(10*sinAng, stairsYDiff)))

  l.AddFeatureByTypeAndPosition(feature.FeatureGuidestone,
    l.upStairsX+int(math.Copysign(15*cosAng, stairsXDiff)),
    l.upStairsY+int(math.Copysign(15*sinAng, stairsYDiff)))
}

func (l *Level) placeCampsites(count int) {
  centerX := l.width / 2
  centerY := l.height / 2

  for i := 0; i < count; i++ {
    ang := ((math.Pi * 2.0) / float64(count)) * (float64(i) + (rand.Float64() / 5.0))
    cosAng := math.Cos(ang)
    sinAng := math.Sin(ang)

    distanceFromCenter :=
      float64(util.Min(l.width, l.height)/2) * (0.5 + (rand.Float64() / 5.0))

    campsiteX := centerX + int(cosAng*distanceFromCenter)
    campsiteY := centerY + int(sinAng*distanceFromCenter)

    l.AddFeatureByTypeAndPosition(feature.FeatureCampsiteSpit, campsiteX, campsiteY)
    l.AddFeatureByTypeAndPosition(feature.FeatureCampsiteBedroll, campsiteX+1, campsiteY)
    l.AddFeatureByTypeAndPosition(feature.FeatureCampsiteBedroll, campsiteX-1, campsiteY)
    l.AddFeatureByTypeAndPosition(feature.FeatureCampsiteBedroll, campsiteX, campsiteY+1)
    l.AddFeatureByTypeAndPosition(feature.FeatureCampsiteBedroll, campsiteX, campsiteY-1)
    l.clearRadius(campsiteX, campsiteY, 5+rand.Intn(2))
  }
}

func (l *Level) clearRadius(x, y, radius int) {
  if radius == 0 {
    return
  }

  halfPi := (math.Pi / 2.0)

  for row := y - radius; row < y+radius; row++ {
    yUnit := math.Sin(halfPi * (1.0 - (math.Abs(float64(y-row)) / float64(radius))))
    rowWidth := int(float64(radius) * 2.0 * yUnit)
    halfRowWidth := rowWidth / 2
    for xCell := x - rowWidth/2; xCell < x+halfRowWidth; xCell++ {
      if IsWall(l.GetTile(xCell, row)) {
        l.SetTile(xCell, row, FreeTile)
      }
    }
  }
}
