package main

import (
	"./command"
	"./interop"
	"./universe"
	"./util"
	"flag"
	"github.com/nsf/termbox-go"
	"hash/crc64"
)

func main() {
	err := termbox.Init()

	if err != nil {
		panic(err)
	}

	defer termbox.Close()
	termbox.HideCursor()

	seed := uint64(1)
	seedFlag := flag.String("seed",
		"default",
		"a word or phrase to seed the dungeon")
	startLevel := uint64(1)
	startLevelFlag := flag.Int("level",
		1,
		"the starting level")
	flag.Parse()

	if seedFlag != nil {
		seed = crc64.Checksum([]byte(*seedFlag), crc64.MakeTable(crc64.ISO))
	}

	if startLevelFlag != nil {
		startLevel = uint64(*startLevelFlag)
	}

	gameLoop(seed, startLevel)
}

func gameLoop(seed, startLevel uint64) {
	u := universe.Universe{}
	u.SetSeed(seed)
	u.GoToLevel(startLevel)

	com := command.Command{}

	displayEverything(&u, &com)

	var endGame bool

	for {
		ev := termbox.PollEvent()
		com.Accept(&ev)

		if com.IsReady() {
			endGame = updateUniverse(&u, &com)
			com.Processed()
		}

		displayEverything(&u, &com)

		if endGame == true {
			break
		}
	}
}

func updateUniverse(u interop.Universe, com interop.Command) bool {
	return u.Advance(com)
}

func displayEverything(u interop.Universe, com interop.Command) {
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	displayUniverse(u)
	displayCommand(com)
	termbox.Flush()
}

func displayUniverse(u interop.Universe) {
	width, height := termbox.Size()
	u.Report(width, height)
}

func displayCommand(com interop.Command) {
	_, height := termbox.Size()
	util.StringOut(0, height-1, com.Format())
}
