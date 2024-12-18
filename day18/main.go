package main

import (
	"fmt"
	"sknoslo/aoc2024/algo"
	"sknoslo/aoc2024/grid"
	"sknoslo/aoc2024/utils"
	"sknoslo/aoc2024/vec2"
	"strconv"
	"strings"

	"github.com/hashicorp/go-set/v3"
)

var input string

func init() {
	input = utils.MustReadInput("input.txt")
}

func main() {
	utils.Run(1, partone)
	utils.Run(2, parttwo)
}

func parseInput() []vec2.Vec2 {
	lines := strings.Split(input, "\n")
	out := make([]vec2.Vec2, len(lines))

	for i, line := range lines {
		nums := utils.MustSplitIntsSep(line, ",")
		out[i] = vec2.New(nums[0], nums[1])
	}

	return out
}

type step struct {
	p vec2.Vec2
	d int
}

func partone() string {
	size := 70
	sim := 1024
	end := vec2.New(size, size)
	bytes := parseInput()
	g := grid.FromSize(end.X+1, end.Y+1, '.')

	for i := range sim {
		g.SetCellAt(bytes[i], '#')
	}

	q := algo.NewDeque[step](71 * 71)
	s := set.New[vec2.Vec2](71 * 71)

	q.PushFront(step{vec2.New(0, 0), 0})

	for !q.Empty() {
		x := q.PopBack()

		if x.p == end {
			return strconv.Itoa(x.d)
		}

		if !g.InGrid(x.p) || g.CellAt(x.p) == '#' || s.Contains(x.p) {
			continue
		}
		s.Insert(x.p)

		for _, dir := range vec2.CardinalDirs {
			q.PushFront(step{x.p.Add(dir), x.d + 1})
		}
	}

	return "no solution"
}

func parttwo() string {
	size := 70
	sim := 1024
	end := vec2.New(size, size)
	bytes := parseInput()
	g := grid.FromSize(end.X+1, end.Y+1, '.')

mainloop:
	for i := range bytes {
		g.SetCellAt(bytes[i], '#')
		// jump the first sim bytes since we know they don't obstruct the path
		if i < sim {
			continue
		}

		q := algo.NewDeque[step](71 * 71)
		s := set.New[vec2.Vec2](71 * 71)

		q.PushFront(step{vec2.New(0, 0), 0})

		for !q.Empty() {
			x := q.PopBack()

			if x.p == end {
				continue mainloop
			}

			if !g.InGrid(x.p) || g.CellAt(x.p) == '#' || s.Contains(x.p) {
				continue
			}
			s.Insert(x.p)

			for _, dir := range vec2.CardinalDirs {
				q.PushFront(step{x.p.Add(dir), x.d + 1})
			}
		}

		return fmt.Sprintf("%d,%d", bytes[i].X, bytes[i].Y)
	}

	return "no solution"
}