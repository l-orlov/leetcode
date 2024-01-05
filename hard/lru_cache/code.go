package lru_cache

type LRUCache struct {
	Head, Tail *Node
	Mp         map[int]*Node
	Capacity   int
}

func newLRUCache(head, tail *Node, cap int) LRUCache {
	return LRUCache{
		Head:     head,
		Tail:     tail,
		Mp:       make(map[int]*Node),
		Capacity: cap,
	}
}

type Node struct {
	Prev, Next *Node
	Key, Value int
}

func newNode(key, val int) *Node {
	return &Node{
		Key:   key,
		Value: val,
	}
}

func Constructor(capacity int) LRUCache {
	head, tail := newNode(0, 0), newNode(0, 0)

	head.Next = tail
	tail.Prev = head
	return newLRUCache(head, tail, capacity)
}

func (c *LRUCache) Get(key int) int {
	if n, ok := c.Mp[key]; ok {
		c.remove(n)
		c.insert(n)
		return n.Value
	}

	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if n, ok := c.Mp[key]; ok {
		c.remove(n)
		n.Value = value
		c.insert(n)
		return
	}

	if len(c.Mp) == c.Capacity {
		delete(c.Mp, c.Tail.Prev.Key)
		c.remove(c.Tail.Prev)
	}

	node := newNode(key, value)
	c.Mp[node.Key] = node
	c.insert(node)
}

func (c *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (c *LRUCache) insert(node *Node) {
	next := c.Head.Next
	c.Head.Next = node
	node.Prev = c.Head
	next.Prev = node
	node.Next = next
}
