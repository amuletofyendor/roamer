package history

type HistoryItem struct {
	history      string
	tickStamp    int
	subTickStamp int
}

func (h *HistoryItem) History() string {
	return h.history
}

func (h *HistoryItem) TickStamp() (int, int) {
	return h.tickStamp, h.subTickStamp
}
