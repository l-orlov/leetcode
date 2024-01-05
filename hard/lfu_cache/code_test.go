package lfu_cache

import (
	"fmt"
	"testing"
)

func TestLFUCache_Get(t *testing.T) {
	const capacity = 5
	obj := Constructor(capacity)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
}
