package level

func (l *Level) CurrentTick() int {
  return l.currentTick
}

func (l *Level) CurrentSubTick() int {
  return l.currentSubTick
}

func (l *Level) SetCurrentTick(t int) {
  l.currentTick = t
}

func (l *Level) SetCurrentSubTick(st int) {
  l.currentSubTick = st
}
