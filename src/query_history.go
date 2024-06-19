package main

type queryHistoryType struct {
	history []string
	index   int
}

var queryHistory queryHistoryType

func (queryHistory *queryHistoryType) back() string {
	if queryHistory.getUpperIndex() > 0 && queryHistory.index > 0 {
		queryHistory.index--
		return queryHistory.history[queryHistory.index]
	}
	queryHistory.index = 0
	return ""
}

func (queryHistory *queryHistoryType) forward() string {
	if queryHistory.index < queryHistory.getUpperIndex() {
		queryHistory.index++
		return queryHistory.history[queryHistory.index]
	}
	queryHistory.resetIndex()
	return ""
}

func (queryHistory *queryHistoryType) add(historyEntry string) {
	queryHistory.history = append(queryHistory.history, historyEntry)
}

func (queryHistory *queryHistoryType) getUpperIndex() int {
	return len(queryHistory.history) - 1
}

func (queryHistory *queryHistoryType) resetIndex() {
	queryHistory.index = queryHistory.getUpperIndex()
}
