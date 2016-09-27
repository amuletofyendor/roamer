package behavior

import (
	"../../interactor"
	"../../interop"
	"../../dice"
)

type FollowBehavior struct {
}

// Behavior ideas
// Dumb, semi-random walk until within line-of-sight
// Pathfinding
// Flanking pathfinding
// Stalking (staying a certain distance until player incapacitated)
// Distractions: food, jewels, a dummy of the player, etc
// Camoflage: dress as a goblin, etc

func (FollowBehavior) Tick(e interop.Entity, com interop.Command, l interop.Level) {
	if dice.Oned8() == 8 ||
		 l.Shrouded(e.X(), e.Y()) {
		return
	}

	var xOff int = e.X() - l.Player().X()
	var yOff int = e.Y() - l.Player().Y()

	wantToMoveLeft := false
	wantToMoveRight := false
	wantToMoveUp := false
	wantToMoveDown := false

	if xOff > 0 {
		wantToMoveLeft = true
	}

	if xOff < 0 {
		wantToMoveRight = true
	}

	if yOff > 0 {
		wantToMoveUp = true
	}

	if yOff < 0 {
		wantToMoveDown = true
	}

	if dice.Oned8() == 1 {
		wantToMoveLeft = !wantToMoveLeft
		wantToMoveRight = !wantToMoveLeft
	}

	if dice.Oned8() == 1 {
		wantToMoveUp = !wantToMoveUp
		wantToMoveDown = !wantToMoveUp
	}

	li := interactor.LevelInteractor{l, e}
	li.Move(wantToMoveLeft, wantToMoveRight, wantToMoveUp, wantToMoveDown)
}
