package pqueues

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := New[rune](3)

	pq.Push('C', 3)
	pq.Push('D', 4)
	pq.Push('A', 1)
	pq.Push('E', 5)
	pq.Push('I', 9)
	pq.Push('H', 8)
	pq.Push('G', 7)
	pq.Push('B', 2)
	pq.Push('F', 6)

	want := "ABCDEFGHI"

	for _, v := range want {
		actual := pq.Pop()
		if v != actual {
			t.Fatalf(`Expected "%v" but got "%v"`, v, actual)
		}
	}
}
