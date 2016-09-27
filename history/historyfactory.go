package history

type HistoryFactory struct {
}

func (hf *HistoryFactory) MakeHistory(history string, tickStamp, subTickStamp int) HistoryItem {
	return HistoryItem{history, tickStamp, subTickStamp}
}
