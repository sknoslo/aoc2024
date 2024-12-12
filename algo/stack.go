package algo

type Stack[T any] struct {
	data []T
	end int
}

func NewStack[T any](size int) *Stack[T] {
	s := new(Stack[T])
	s.data = make([]T, size)
	s.end = -1
	return s
}

func (stack *Stack[T]) Push(item T) {
	stack.end++
	stack.grow()
	stack.data[stack.end] = item
}

func (stack *Stack[T]) Empty() bool {
	return stack.end < 0
}

func (stack *Stack[T]) Pop() T {
	var empty T
	res := stack.data[stack.end]
	stack.data[stack.end] = empty
	stack.end--
	return res
}

func (stack *Stack[T]) Clear() {
	clear(stack.data)
	stack.end = -1
}

func (stack *Stack[T]) grow() {
	if stack.end < len(stack.data) {
		return
	}

	tmp := stack.data
	stack.data = make([]T, len(stack.data) * 2)
	copy(stack.data, tmp)
}
