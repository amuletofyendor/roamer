package util

import (
	"github.com/nsf/termbox-go"
)

func StringOut(x, y int, s string) {
	for i, c := range s {
		termbox.SetCell(x+i, y, c, termbox.ColorWhite, termbox.ColorBlack)
	}
}
