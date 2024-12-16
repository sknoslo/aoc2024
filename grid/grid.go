package grid

import (
	"fmt"
	"iter"
	"sknoslo/aoc2024/utils"
	"sknoslo/aoc2024/vec2"
	"strings"
)

type Grid[T comparable] struct {
	w, h int
	cells []T
}

func MustFromDigits(in string) *Grid[int] {
	lines := strings.Split(in, "\n")
	w, h := len(lines[0]), len(lines)
	cells := make([]int, w * h)

	for i, v := range strings.Join(lines, "") {
		cells[i] = utils.MustAtoi(string(v))
	}

	return New(w, h, cells)
}

func FromRunes(in string) *Grid[rune] {
	lines := strings.Split(in, "\n")
	w, h := len(lines[0]), len(lines)
	cells := make([]rune, w * h)

	for i, v := range strings.Join(lines, "") {
		cells[i] = v
	}

	return New(w, h, cells)
}

func New[T comparable](w, h int, cells []T) *Grid[T] {
	return &Grid[T]{w, h, cells}
}

func (grid *Grid[T]) CellAt(v vec2.Vec2) T {
	i := v.Y * grid.w + v.X
	return grid.cells[i]
}

func (grid *Grid[T]) CellAtXY(x, y int) T {
	i := y * grid.w + x
	return grid.cells[i]
}

func (grid *Grid[T]) SetCellAt(v vec2.Vec2, t T) {
	i := v.Y * grid.w + v.X
	grid.cells[i] = t
}

func (grid *Grid[T]) InGrid(v vec2.Vec2) bool {
	return v.InRange(0, 0, grid.w-1, grid.h-1)
}

func (grid *Grid[T]) Find(item T) vec2.Vec2 {
	for i, v := range grid.cells {
		if v == item {
			return grid.indexToVec2(i)
		}
    }
	return vec2.New(0, 0)
}

func (grid *Grid[T]) Cells() iter.Seq2[vec2.Vec2, T] {
	return func(yield func(vec2.Vec2, T) bool) {
		for i, v := range grid.cells {
			if !yield(grid.indexToVec2(i), v) {
				return
			}
		}
	}
}

func (grid *Grid[T]) Points() iter.Seq[vec2.Vec2] {
	return func(yield func(vec2.Vec2) bool) {
		for i := range grid.cells {
			if !yield(grid.indexToVec2(i)) {
				return
			}
		}
	}
}

func (grid *Grid[T]) String() string {
	var b strings.Builder

	for s, v := range grid.Cells() {
		fmt.Fprint(&b, v)
		if s.X == grid.w - 1 {
			fmt.Fprintln(&b)
		}
	}

	return b.String()
}

func (grid *Grid[T]) Stringf(f string) string {
	var b strings.Builder

	for s, v := range grid.Cells() {
		fmt.Fprintf(&b, f, v)
		if s.X == grid.w - 1 {
			fmt.Fprintln(&b)
		}
	}

	return b.String()
}

func (grid *Grid[T]) StringOverlayf(f string, overlay T, at vec2.Vec2) string {
	var b strings.Builder

	for s, v := range grid.Cells() {
		if s == at {
			v = overlay
		}
		fmt.Fprintf(&b, f, v)
		if s.X == grid.w - 1 {
			fmt.Fprintln(&b)
		}
	}

	return b.String()
}

func (grid *Grid[T]) StringOverlayMapf(f string, overlay T, m map[vec2.Vec2]int) string {
	var b strings.Builder

	for s, v := range grid.Cells() {
		if _, ok := m[s]; ok {
			v = overlay
		}
		fmt.Fprintf(&b, f, v)
		if s.X == grid.w - 1 {
			fmt.Fprintln(&b)
		}
	}

	return b.String()
}

func (grid *Grid[T]) indexToVec2(i int) vec2.Vec2 {
	return vec2.New(i % grid.w, i / grid.w)
}
