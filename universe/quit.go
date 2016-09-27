package universe

import (
  "../interop"
)

func (u *Universe) WantsQuit() bool {
  return u.quitLevel == 1
}

func (u *Universe) ReallyQuit() bool {
  return u.quitLevel >= 1
}

func (u *Universe) AdvanceQuit() {
  u.quitLevel += 1
}

func (u *Universe) ResetQuit() {
  u.quitLevel = 0
}

func (u *Universe) HandleQuitting(com interop.Command) {
  if com.IsQuit() {
    com.ChoiceMode([]rune{'y', 'n'}, "quit choice", "Are you sure you want to quit?")
  } else if com.IsQuitChoice() && com.Choice() == 'y' {
    u.AdvanceQuit()
  } else {
    u.ResetQuit()
  }
}
