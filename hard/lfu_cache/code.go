package lfu_cache

type (
	LFUCache struct {
		capacity   int
		elements   map[int]*ListElement
		freqGroups map[int]*List
	}
)

func Constructor(capacity int) LFUCache {
	elements := make(map[int]*ListElement, capacity)
	freqGroups := make(map[int]*List)

	return LFUCache{
		capacity:   capacity,
		elements:   elements,
		freqGroups: freqGroups,
	}
}

func (c *LFUCache) Get(key int) int {
	element, ok := c.elements[key]
	if !ok {
		return -1
	}

	// Обновляем частоту
	c.update(key, element.val)

	return element.val
}

func (c *LFUCache) Put(key int, value int) {
	_, ok := c.elements[key]
	if ok {
		// Обновляем частоту и значение
		c.update(key, value)
	}

	if len(c.elements) < c.capacity {
		// Добавляем новый элемент
		c.addNew(key, value)
	}

	// Удаляем least frequently used элемент
	// evict
	// Добавляем новый элемент
	c.addNew(key, value)
}

func (c *LFUCache) addNew(key int, value int) {
	list := c.freqGroups[1]
	if list == nil {
		list = NewList()
	}

	element := &ListElement{
		val:         value,
		freqCounter: 1,
		prev:        nil,
		next:        nil,
	}

	list.AddLast(element)

	c.elements[key] = element
	c.freqGroups[1] = list
}

func (c *LFUCache) update(key int, value int) {
	element := c.elements[key]

	freq := element.freqCounter
	prev := element.prev
	next := element.next

	// Добавляем элемент для новой частоты
	newList := c.freqGroups[freq+1]
	if newList == nil {
		newList = NewList()
	}

	element.val = value
	element.freqCounter++
	element.prev = nil
	element.next = nil
	newList.AddLast(element)

	// Удаляем элемент для прошлой частоты
	list := c.freqGroups[freq]
	if prev == nil {
		// Обновляем первый элемент в списке
		if next == nil {
			// Больше нет элементов => удаляем список
			c.freqGroups[freq] = nil
			return
		}

		// Есть еще элементы => меняем ссылки
		prev.next = next
		list.first = next
		return
	}

	if next == nil {
		// Обновляем последний элемент в списке
		prev.next = nil
		list.last = prev
		return
	}

	// Обновляем элемент в середине списка => меняем ссылки
	prev.next = next
}

func (c *LFUCache) evict(key int, value int) {
	list := c.freqGroups[1]
	if list == nil {
		list = NewList()
	}

	element := &ListElement{
		val:         value,
		freqCounter: 1,
		prev:        nil,
		next:        nil,
	}

	list.AddLast(element)

	c.elements[key] = element
	c.freqGroups[1] = list
}
