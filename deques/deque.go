package deques


/**
 * start points to front of queue
 * end points to end of queue + 1
 * if queue is empty start == end
 */
type Deque[T any] struct {
	data []T
	start, end int
}

func New[T any](size int) *Deque[T] {
	d := new(Deque[T])
	d.data = make([]T, size)
	d.start = size / 2
	d.end = size / 2
	return d
}

// [0 0 0 0 0 0]
//        *
// [0 0 A 0 0 0]
//      s e

func (d *Deque[T]) PushFront(val T) {
	d.grow()
	d.start--
	if d.start < 0 {
		d.start = len(d.data) - 1
	}
	d.data[d.start] = val
}

// [0 0 0 0 0 0]
//        *
// [0 0 0 A 0 0]
//        s e

func (d *Deque[T]) PushBack(val T) {
	d.grow()
	d.data[d.end] = val
	d.end++
	if d.end == len(d.data) {
		d.end = 0
	}
}

func (d *Deque[T]) PopFront() T {
	var empty T
	res := d.data[d.start]
	d.data[d.start] = empty
	d.start++
	if d.start == len(d.data) {
		d.start = 0
	}
	return res
}

func (d *Deque[T]) PopBack() T {
	var empty T
	d.end--
	if d.end < 0 {
		d.end = len(d.data) - 1
	}
	res := d.data[d.end]
	d.data[d.end] = empty
	return res
}

func (d *Deque[T]) Empty() bool {
	return d.start == d.end
}

func (d *Deque[T]) grow() {
	size := len(d.data)
	if d.start == d.end + 1 || d.start == 0 && d.end == size {
		data := make([]T, size * 2)

		s := d.start
		d.start = size / 2 + 1 // start at a quarter
		e := d.start
		for s != d.end {
			data[e] = d.data[s]
			s = (s + 1) % size
			e++
		}
		d.data = data
		d.end = e
	}
}
