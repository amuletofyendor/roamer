package interop

type HistoryItem interface {
	History() string
	TickStamp() (int, int)
}

type HistoryProvider interface {
	HistoryAvailable() bool
	HistoryViewed()
	History() []HistoryItem
	AppendHistory(event string, currentTick, currentSubTick int)
}
