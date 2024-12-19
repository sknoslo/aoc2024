package deques

import (
	"fmt"
	"testing"
)

func TestQueueFrontBack(t *testing.T) {
	q := New[int](6)

	q.PushFront(1)
	q.PushFront(2)
	q.PushFront(3)

	want := []int{1, 2, 3}

	for _, v := range want {
		actual := q.PopBack()

		if actual != v {
			t.Fatalf(`Expected "%v" but got "%v"`, v, actual)
		}
	}
	if !q.Empty() {
		t.Fatal("Expected q to be empty")
	}
}

func TestQueueBackFront(t *testing.T) {
	q := New[int](6)

	q.PushBack(1)
	q.PushBack(2)
	q.PushBack(3)

	want := []int{1, 2, 3}

	for _, v := range want {
		actual := q.PopFront()

		if actual != v {
			t.Fatalf(`Expected "%v" but got "%v"`, v, actual)
		}
	}
	if !q.Empty() {
		t.Fatal("Expected q to be empty")
	}
}

func TestResize(t *testing.T) {
	q := New[int](4)

	for i := range 8 {
		q.PushFront(i)

		fmt.Println(q.data)
		fmt.Println(q.start, q.end)
	}

	fmt.Println("POPPING")

	for i := range 8 {
		actual := q.PopBack()

		fmt.Println(q.data)
		fmt.Println(q.start, q.end)

		if actual != i {
			t.Fatalf(`Expected "%v" but got "%v"`, i, actual)
		}
	}
}
