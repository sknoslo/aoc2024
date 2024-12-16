package algo

import (
	"container/heap"
)

type item[T any] struct {
	value T
	priority int
	index int
}

type priorityQueueData[T any] []item[T]


func (pq priorityQueueData[T]) Len() int {
	return len(pq)
}

func (pq priorityQueueData[T]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq priorityQueueData[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueueData[T]) Push(i any) {
	n := len(*pq)
	it := i.(item[T])
	it.index = n
	*pq = append(*pq, it)
}

func (pq *priorityQueueData[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0:n-1]

	return item
}

type PriorityQueue[T any] struct {
	data priorityQueueData[T]
}

func NewPriorityQueue[T any](size int) *PriorityQueue[T] {
	pq := new(PriorityQueue[T])
	pq.data = make(priorityQueueData[T], 0, size)

	heap.Init(&pq.data)

	return pq
}

func (pq *PriorityQueue[T]) Push(value T, priority int) {
	heap.Push(&pq.data, item[T]{ value, priority, -1 })
}

func (pq *PriorityQueue[T]) Pop() T {
	return heap.Pop(&pq.data).(item[T]).value
}

func (pq *PriorityQueue[T]) Empty() bool {
	return len(pq.data) == 0
}

func (pq *PriorityQueue[T]) Peek() T {
	return pq.data[len(pq.data)-1].value
}

