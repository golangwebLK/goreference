package heap

import (
	"log"
	"testing"
)

func TestHeapInt(t *testing.T) {
	h := NewHeap([]int{2, 1, 3}, func(data []int, i, j int) bool {
		return i < data[j]
	})
	h.Push(4)
	h.Push(5)
	h.Push(6)
	h.Push(7)
	h.Push(8)
	h.Push(9)
	for {
		if item, ok := h.Pop(); ok {
			log.Println(item)
		} else {
			break
		}
	}
}
func TestHeapNode(t *testing.T) {
	type Node struct {
		Value int
		Str   string
	}

	h := NewHeap([]Node{{1, "1"}, {2, "2"}, {3, "3"}}, func(data []Node, newData Node, j int) bool {
		return newData.Value < data[j].Value
	})
	h.Push(Node{4, "4"})
	h.Push(Node{5, "5"})
	h.Push(Node{6, "6"}) // 添加节点 {6, "6"}
	h.Push(Node{7, "7"}) // 添加节点 {7, "7"}
	for {
		if item, ok := h.Pop(); ok {
			log.Println(item)
		} else {
			break
		}
	}

}
