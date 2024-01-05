package freq_stack

type NumList struct {
	num  int
	prev *NumList
}

func (list *NumList) pushTop(num int) *NumList {
	return &NumList{
		num:  num,
		prev: list,
	}
}

type FreqStack struct {
	freq    map[int]int
	maxFreq int
	byFreq  map[int]*NumList
}

func Constructor() FreqStack {
	return FreqStack{
		freq:   make(map[int]int),
		byFreq: make(map[int]*NumList),
	}
}

func (f *FreqStack) Push(val int) {
	freq := f.freq[val] + 1
	if freq > f.maxFreq { // updating maxFreq
		f.maxFreq = freq
	}
	list := f.byFreq[freq]
	f.byFreq[freq] = list.pushTop(val) // adding val to the end of the list
	f.freq[val] = freq
}

func (f *FreqStack) Pop() int {
	list := f.byFreq[f.maxFreq]
	if list == nil {
		return -1
	}

	val := list.num                 // val is the last element of the list
	f.byFreq[f.maxFreq] = list.prev // removing last element from the list
	if f.byFreq[f.maxFreq] == nil { // if no more numbers with frequency maxFreq we decrease maxFreq
		f.maxFreq--
	}
	f.freq[val]-- // decreasing frequency of val
	return val
}
