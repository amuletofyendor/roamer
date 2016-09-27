package level

import (
  "../creature"
  "../interop"
  "../util"
  "../history"
  "fmt"
  "github.com/nsf/termbox-go"
  "sort"
)

type Level struct {
  seed             uint64
  depth            uint64
  currentTick      int
  currentSubTick   int
  layout           [][]rune
  shroud           [][]byte
  entities         []interop.Entity
  entityRenderList []interop.WorldObject
  history          []interop.HistoryItem
  width            int
  height           int
  downStairsX      int
  downStairsY      int
  upStairsX        int
  upStairsY        int
}

func (l *Level) Init(universeSeed, depth uint64) {
  l.seed = universeSeed
  l.depth = depth
  l.initGeography()
  l.populate()
}

func (l *Level) Width() int {
  return l.width
}

func (l *Level) Height() int {
  return l.height
}

func (l *Level) SetTile(x, y int, tile rune) {
  if x >= 0 && y >= 0 && x < l.width && y < l.height {
    l.layout[y][x] = tile
  }
}

func (l *Level) GetTile(x, y int) rune {
  if x >= 0 && y >= 0 && x < l.width && y < l.height {
    return l.layout[y][x]
  } else {
    return InaccessibleTile
  }
}

func (l *Level) Tick(com interop.Command) {
  l.currentTick++
  l.currentSubTick = 0
  player := l.Player()

  for i, _ := range l.entities {
    l.currentSubTick += 1000
    l.entities[i].Tick(com, l)
  }

  l.tidyUp()

  l.deshroudByRay(player.X(), player.Y())
  l.shroudTick()

  hf := history.HistoryFactory{}

  for i := len(l.entities) - 1; i >= 0; i-- {
    e := l.entities[i]

    if e.HistoryAvailable() {
      entityHistory := e.History()
      if entityHistory != nil {
        for _, h := range entityHistory {
          tickStamp, subTickStamp := h.TickStamp()
          historyItem := hf.MakeHistory(fmt.Sprintf("%v %v", e.Name(), h.History()),
                           tickStamp,
                           subTickStamp)
          l.history = append(l.history, &historyItem)
        }

        e.HistoryViewed()
      }
    }

    if !e.IsPlayer() && !e.IsAlive() {
      eL, ok := e.(*creature.Creature)
      if ok == true {
        l.Corpseify(eL)
      }
    }
  }

  sort.Sort(history.HistoryItemByTick(l.history))
}

func (l *Level) Report(xStart, yStart, maxX, maxY int) {
  worldOffX := l.Player().X() - (maxX / 2)
  worldOffY := l.Player().Y() - (maxY / 2)

  var dispX, dispY, tileX, tileY int
  var tile rune

  for y := 0; y < maxY; y++ {
    for x := 0; x < maxX; x++ {
      tileX = x + worldOffX
      tileY = y + worldOffY
      dispX = xStart + x
      dispY = yStart + y

      if l.Shrouded(tileX, tileY) {
        tile = '?'
      } else {
        tile = l.GetTile(tileX, tileY)
      }

      if dispX >= xStart && dispY >= yStart &&
        dispX < maxX && dispY < maxY {
        termbox.SetCell(dispX,
          dispY,
          tile,
          termbox.ColorWhite,
          termbox.ColorBlack)
      }
    }
  }

  for _, e := range *l.renderList() {
    dispX := xStart + e.X() - worldOffX
    dispY := yStart + e.Y() - worldOffY

    if dispX >= xStart && dispY >= yStart && dispX < maxX && dispY < maxY {
      if l.Shrouded(e.X(), e.Y()) {
        tile = '?'
      } else {
        tile = e.DisplayRune()
      }

      termbox.SetCell(dispX,
        dispY,
        tile,
        termbox.ColorWhite,
        termbox.ColorBlack)
    }
  }

  for i, h := range l.history {
    util.StringOut(maxX/2, i+1, fmt.Sprintf("> %v", h.History()))
  }

  l.history = nil
}
