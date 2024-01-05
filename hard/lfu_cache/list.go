package lfu_cache

type (
	List struct {
		first *ListElement
		last  *ListElement
	}
	ListElement struct {
		val         int
		freqCounter int
		prev        *ListElement
		next        *ListElement
	}
)

func NewList() *List {
	return &List{}
}

func (l *List) AddLast(element *ListElement) {
	if l == nil {
		return
	}

	if l.last != nil {
		l.last.next = element
		element.prev = l.last
		return
	}

	l.first = element
	l.last = element
}
