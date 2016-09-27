package behavior

import (
	"../../interactor"
	"../../interop"
)

type PlayerBehavior struct {
}

func (PlayerBehavior) Tick(e interop.Entity, com interop.Command, l interop.Level) {
	li := interactor.LevelInteractor{l, e}

	if com.IsMove() {
		wantToMoveLeft := false
		wantToMoveRight := false
		wantToMoveUp := false
		wantToMoveDown := false

		if com.IsLeft() {
			wantToMoveLeft = true
		}

		if com.IsRight() {
			wantToMoveRight = true
		}

		if com.IsUp() {
			wantToMoveUp = true
		}

		if com.IsDown() {
			wantToMoveDown = true
		}

		li.Move(wantToMoveLeft, wantToMoveRight, wantToMoveUp, wantToMoveDown)
	} else if com.IsLookHere() {
		li.Look(e.X(), e.Y())
	} else if com.IsTake() {
		li.Take(e.X(), e.Y())
	} else if com.IsWield() {
		weapons := e.NameMap("weapons")
		weaponKeys := make([]rune, 0, len(weapons))

		for k, _ := range weapons {
			weaponKeys = append(weaponKeys, k)
		}

		if len(weaponKeys) > 0 {
			com.ChoiceMode(weaponKeys, "wield choice", "Choose a weapon to wield")
		}
	} else if com.IsWieldChoice() {
		li.Wield(com.Choice())
	} else if com.IsEat() {
		edibles := e.NameMap("edibles")
		edibleKeys := make([]rune, 0, len(edibles))

		for k, _ := range edibles {
			edibleKeys = append(edibleKeys, k)
		}

		if len(edibleKeys) > 0 {
			com.ChoiceMode(edibleKeys, "eating choice", "Choose something to eat")
		}
	} else if com.IsEatingChoice() {
		li.Eat(com.Choice())
	} else if com.IsCraft() {
	}
}
