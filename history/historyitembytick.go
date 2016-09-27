package history

import(
  "../interop"
)

type HistoryItemByTick []interop.HistoryItem

func (a HistoryItemByTick) Len() int {
	return len(a)
}

func (a HistoryItemByTick) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a HistoryItemByTick) Less(i, j int) bool {
  tickStampi, subTickStampi := a[i].TickStamp()
  tickStampj, subTickStampj := a[j].TickStamp()

	return (tickStampi < tickStampj) ||
		     ((tickStampi == tickStampj) &&
		      (subTickStampi < subTickStampj))
}
