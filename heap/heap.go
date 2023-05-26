package heap

import "reflect"

type Heap[T any] struct {
	lessCompare LessCompare[T]
	data        []T
}
type LessCompare[T any] func(data []T, newData T, j int) bool

func NewHeap[T any](data []T, lessCompare LessCompare[T]) *Heap[T] {
	h := &Heap[T]{
		lessCompare: lessCompare,
		data:        data,
	}
	h.init()
	return h
}

func (h *Heap[T]) init() {
	n := len(h.data)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}
}

func (h *Heap[T]) Push(data T) {
	h.data = append(h.data, data)
	h.up(len(h.data) - 1)
}

func (h *Heap[T]) up(index int) {
	endData := h.data[index]
	for h.lessCompare(h.data, endData, (index-1)/2) {
		h.data[index] = h.data[(index-1)/2]
		index = (index - 1) / 2
		if index == 0 {
			h.data[index] = endData
			return
		}
	}
	h.data[index] = endData
}

func (h *Heap[T]) Pop() (T, bool) {
	if len(h.data) == 0 {
		return reflect.Zero(reflect.TypeOf((*T)(nil)).Elem()).Interface().(T), false
	}
	item := h.data[0]
	lastIndex := len(h.data) - 1
	h.data[0] = h.data[lastIndex]
	h.data = h.data[:lastIndex]
	if len(h.data) > 0 {
		h.down(0)
	}
	return item, true
}

func (h *Heap[T]) down(index int) {
	firstData := h.data[index]
	for {
		leftChildIndex := 2*index + 1
		rightChildIndex := leftChildIndex + 1

		if leftChildIndex >= len(h.data) {
			break
		}
		childIndex := leftChildIndex

		if rightChildIndex < len(h.data) && h.lessCompare(h.data, h.data[rightChildIndex], leftChildIndex) {
			childIndex = rightChildIndex
		}
		if h.lessCompare(h.data, firstData, childIndex) {
			break
		}
		h.data[index] = h.data[childIndex]
		index = childIndex
	}
	h.data[index] = firstData
}
